package main

import (
	"context"
	"log"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	spotify "github.com/ericflores108/samba/pkg/spotify"
)

const (
	listView uint = iota
	titleView
	bodyView
)

type model struct {
	state     uint
	user      *spotify.PrivateUserObject
	client    *spotify.APIClient
	ctx       context.Context
	tracks    []spotify.SavedTrackObject
	currTrack spotify.SavedTrackObject
	listIndex int
	textarea  textarea.Model
	textinput textinput.Model
}

func NewModel(user *spotify.PrivateUserObject, client *spotify.APIClient, ctx context.Context) model {
	// Get user's saved tracks
	tracksResp, _, err := client.LibraryAPI.GetUsersSavedTracks(ctx).Execute()
	if err != nil {
		log.Fatalf("Failed to get saved tracks: %v", err)
	}

	tracks := []spotify.SavedTrackObject{}
	if tracksResp != nil && tracksResp.Items != nil {
		tracks = *tracksResp.Items
	}

	return model{
		state:     listView,
		user:      user,
		client:    client,
		ctx:       ctx,
		tracks:    tracks,
		textarea:  textarea.New(),
		textinput: textinput.New(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmds []tea.Cmd
		cmd  tea.Cmd
	)

	m.textarea, cmd = m.textarea.Update(msg)
	cmds = append(cmds, cmd)

	m.textinput, cmd = m.textinput.Update(msg)
	cmds = append(cmds, cmd)

	switch msg := msg.(type) {
	// handle key strokes
	case tea.KeyMsg:
		key := msg.String()
		switch m.state {
		// List View key bindings
		case listView:
			switch key {
			case "q":
				return m, tea.Quit
			case "r":
				// Refresh tracks
				tracksResp, _, err := m.client.LibraryAPI.GetUsersSavedTracks(m.ctx).Execute()
				if err == nil && tracksResp != nil && tracksResp.Items != nil {
					m.tracks = *tracksResp.Items
				}
			case "up", "k":
				if m.listIndex > 0 {
					m.listIndex--
				}
			case "down", "j":
				if m.listIndex < len(m.tracks)-1 {
					m.listIndex++
				}
			case "enter":
				if len(m.tracks) > 0 {
					m.currTrack = m.tracks[m.listIndex]
					m.state = bodyView
					trackInfo := "Track: " + m.currTrack.Track.GetName() + "\n"
					if len(m.currTrack.Track.Artists) > 0 {
						trackInfo += "Artist: " + m.currTrack.Track.Artists[0].GetName() + "\n"
					}
					if m.currTrack.Track.Album != nil {
						trackInfo += "Album: " + m.currTrack.Track.Album.GetName() + "\n"
					}
					m.textarea.SetValue(trackInfo)
					m.textarea.Focus()
					m.textarea.CursorEnd()
				}
			}

		// No title view needed for Spotify tracks

		// Track info view key bindings
		case bodyView:
			switch key {
			case "esc":
				m.state = listView
			}
		}
	}

	return m, tea.Batch(cmds...)
}
