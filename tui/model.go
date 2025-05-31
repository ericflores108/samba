package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	spotify "github.com/ericflores108/samba/pkg/spotify"
	sw "github.com/ericflores108/samba/pkg/spotifywrapper"
)

const (
	authWelcomeView uint = iota
	authURLView
	authCodeView
	authLoadingView
	listView
	titleView
	bodyView
	itemsPerPage = 10
)

type authURLMsg struct {
	URL   string
	State string
}

type authCodeMsg struct {
	Code  string
	State string
}

type authSuccessMsg struct {
	User *spotify.PrivateUserObject
}

type authErrorMsg struct {
	Error string
}

type currentlyPlayingMsg struct {
	CurrentlyPlaying *spotify.CurrentlyPlayingContextObject
}

type model struct {
	state           uint
	user            *spotify.PrivateUserObject
	client          *sw.SpotifyClient
	ctx             context.Context
	listIndex       int
	textarea        textarea.Model
	textinput       textinput.Model
	authURL         string
	authState       string
	authCode        string
	errorMsg        string
	server          *http.Server
	currentlyPlaying *spotify.CurrentlyPlayingContextObject
	*spotify.QueueObject
}

func NewModel(client *sw.SpotifyClient, ctx context.Context) model {
	ti := textinput.New()
	ti.Placeholder = "Enter authorization code here..."
	ti.Focus()
	ti.CharLimit = 500
	ti.Width = 60

	return model{
		state:     authWelcomeView,
		client:    client,
		ctx:       ctx,
		textarea:  textarea.New(),
		textinput: ti,
	}
}

func NewAuthenticatedModel(user *spotify.PrivateUserObject, client *sw.SpotifyClient, ctx context.Context) model {
	// Get user's saved tracks
	q, res, err := client.GetQueue(ctx)
	if err != nil {
		log.Fatalf("client.PlayerAPI.GetQueueExecute: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		log.Fatalf("client.PlayerAPI.GetQueueExecute: %d", res.StatusCode)
	}

	// Get currently playing track
	currentlyPlaying, _, err := client.GetCurrentlyPlaying(ctx)
	if err != nil {
		// Don't fail if no track is currently playing
		log.Printf("Warning: Could not get currently playing track: %v", err)
	}

	return model{
		state:            listView,
		user:             user,
		client:           client,
		ctx:              ctx,
		textarea:         textarea.New(),
		textinput:        textinput.New(),
		currentlyPlaying: currentlyPlaying,
		QueueObject:      q,
	}
}

func (m model) Init() tea.Cmd {
	if m.state == authWelcomeView {
		return generateAuthURL(m.client)
	}
	return nil
}

func generateAuthURL(client *sw.SpotifyClient) tea.Cmd {
	return func() tea.Msg {
		authURL, state, err := client.GetAuthURL()
		if err != nil {
			return authErrorMsg{Error: fmt.Sprintf("Failed to generate auth URL: %v", err)}
		}
		return authURLMsg{URL: authURL, State: state}
	}
}

func startCallbackServer() tea.Cmd {
	return func() tea.Msg {
		listener, err := net.Listen("tcp", ":8080")
		if err != nil {
			return authErrorMsg{Error: fmt.Sprintf("Failed to start callback server: %v", err)}
		}

		codeChan := make(chan string, 1)
		stateChan := make(chan string, 1)

		mux := http.NewServeMux()
		mux.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
			code := r.URL.Query().Get("code")
			state := r.URL.Query().Get("state")

			if code != "" {
				codeChan <- code
				stateChan <- state
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("Authorization successful! You can close this window and return to the terminal."))
			} else {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Authorization failed: no code received"))
			}
		})

		server := &http.Server{Handler: mux}
		go func() {
			server.Serve(listener)
		}()

		// Wait for the code with a timeout
		select {
		case code := <-codeChan:
			state := <-stateChan
			server.Shutdown(context.Background())
			return authCodeMsg{Code: code, State: state}
		case <-time.After(5 * time.Minute):
			server.Shutdown(context.Background())
			return authErrorMsg{Error: "Authorization timeout - please try again"}
		}
	}
}

func exchangeCodeForToken(client *sw.SpotifyClient, code, state string) tea.Cmd {
	return func() tea.Msg {
		if err := client.ExchangeCodeForToken(code, state); err != nil {
			return authErrorMsg{Error: fmt.Sprintf("Failed to exchange code for token: %v", err)}
		}

		// Get user profile to verify authentication
		user, _, err := client.GetCurrentUserProfile(context.Background())
		if err != nil {
			return authErrorMsg{Error: fmt.Sprintf("Failed to get user profile: %v", err)}
		}

		return authSuccessMsg{User: user}
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmds []tea.Cmd
		cmd  tea.Cmd
	)

	switch msg := msg.(type) {
	case authURLMsg:
		m.authURL = msg.URL
		m.authState = msg.State
		m.state = authURLView
		return m, startCallbackServer()

	case authCodeMsg:
		m.authCode = msg.Code
		m.state = authLoadingView
		return m, exchangeCodeForToken(m.client, msg.Code, msg.State)

	case authSuccessMsg:
		m.user = msg.User
		// Get user's queue and transition to main app
		q, res, err := m.client.GetQueue(m.ctx)
		if err != nil {
			m.errorMsg = fmt.Sprintf("Failed to get queue: %v", err)
			return m, nil
		}
		if res.StatusCode != http.StatusOK {
			m.errorMsg = fmt.Sprintf("Failed to get queue: status %d", res.StatusCode)
			return m, nil
		}
		m.QueueObject = q
		
		// Get currently playing track
		currentlyPlaying, _, err := m.client.GetCurrentlyPlaying(m.ctx)
		if err != nil {
			// Don't fail if no track is currently playing
			log.Printf("Warning: Could not get currently playing track: %v", err)
		}
		m.currentlyPlaying = currentlyPlaying
		
		m.state = listView
		return m, nil

	case authErrorMsg:
		m.errorMsg = msg.Error
		return m, nil

	case tea.KeyMsg:
		key := msg.String()
		switch m.state {
		case authWelcomeView:
			switch key {
			case "q", "ctrl+c":
				return m, tea.Quit
			case "enter", " ":
				return m, generateAuthURL(m.client)
			}

		case authURLView:
			switch key {
			case "q", "ctrl+c":
				return m, tea.Quit
			case "r":
				return m, generateAuthURL(m.client)
			}

		case authCodeView:
			m.textinput, cmd = m.textinput.Update(msg)
			cmds = append(cmds, cmd)
			switch key {
			case "q", "ctrl+c":
				return m, tea.Quit
			case "enter":
				if m.textinput.Value() != "" {
					return m, exchangeCodeForToken(m.client, m.textinput.Value(), m.authState)
				}
			case "esc":
				m.state = authURLView
			}

		case authLoadingView:
			switch key {
			case "q", "ctrl+c":
				return m, tea.Quit
			}

		// List View key bindings
		case listView:
			switch key {
			case "q":
				return m, tea.Quit
			case "up", "k":
				if m.listIndex > 0 {
					m.listIndex--
				}
			case "down", "j":
				if m.listIndex < len(m.QueueObject.GetQueue())-1 {
					m.listIndex++
				}
			case "left", "h":
				// Go to previous page
				currentPage := m.listIndex / itemsPerPage
				if currentPage > 0 {
					m.listIndex = (currentPage - 1) * itemsPerPage
				}
			case "right", "l":
				// Go to next page
				totalItems := len(m.QueueObject.GetQueue())
				currentPage := m.listIndex / itemsPerPage
				totalPages := (totalItems + itemsPerPage - 1) / itemsPerPage
				if currentPage < totalPages-1 {
					m.listIndex = (currentPage + 1) * itemsPerPage
					if m.listIndex >= totalItems {
						m.listIndex = totalItems - 1
					}
				}
			}

		// Title Input View key bindings
		case titleView:
			m.textinput, cmd = m.textinput.Update(msg)
			cmds = append(cmds, cmd)
			switch key {
			case "esc":
				m.state = listView
			}

		// Body Textarea key bindings
		case bodyView:
			m.textarea, cmd = m.textarea.Update(msg)
			cmds = append(cmds, cmd)
			switch key {
			case "esc":
				m.state = listView
			}
		}
	}

	return m, tea.Batch(cmds...)
}
