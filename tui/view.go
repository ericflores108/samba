package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

var (
	appNameStyle = lipgloss.NewStyle().Background(lipgloss.Color("99")).Padding(0, 1)

	faint = lipgloss.NewStyle().Foreground(lipgloss.Color("255")).Faint(true)

	listEnumeratorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("99")).MarginRight(1)

	linkStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("12")).Underline(true)

	// Now Playing styles
	nowPlayingBoxStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("99")).
		Padding(1, 2).
		MarginBottom(1).
		Background(lipgloss.Color("235"))

	nowPlayingTitleStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("198")).
		Bold(true)

	nowPlayingArtistStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("99")).
		Italic(true)

	// Enhanced queue styles
	queueHeaderStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("226")).
		Bold(true).
		MarginTop(1).
		MarginBottom(1)

	queueItemStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("255")).
		Background(lipgloss.Color("238")).
		Padding(0, 1).
		MarginBottom(0)

	queueItemSelectedStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("226")).
		Background(lipgloss.Color("240")).
		Padding(0, 1).
		MarginBottom(0).
		Bold(true)

	queueEnumeratorStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("99")).
		MarginRight(1).
		Bold(true)

	queueEnumeratorSelectedStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("226")).
		MarginRight(1).
		Bold(true)
)

// wrapText wraps text to fit within the given width
func wrapText(text string, width int) string {
	if width <= 0 {
		width = 80 // default width
	}

	words := strings.Fields(text)
	if len(words) == 0 {
		return text
	}

	var lines []string
	currentLine := ""

	for _, word := range words {
		if len(currentLine) == 0 {
			currentLine = word
		} else if len(currentLine)+1+len(word) <= width {
			currentLine += " " + word
		} else {
			lines = append(lines, currentLine)
			currentLine = word
		}
	}

	if len(currentLine) > 0 {
		lines = append(lines, currentLine)
	}

	return strings.Join(lines, "\n")
}

func (m model) View() string {
	// Get terminal width for text wrapping
	width := 80 // default width
	if fd := int(os.Stdout.Fd()); term.IsTerminal(fd) {
		if termWidth, _, err := term.GetSize(fd); err == nil {
			width = termWidth - 2 // leave minimal margin
		}
	}

	s := appNameStyle.Render("SAMBA TUI") + "\n\n"

	// Show error message if present
	if m.errorMsg != "" {
		s += lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Render("Error: "+m.errorMsg) + "\n\n"
	}

	switch m.state {
	case authWelcomeView:
		s += "Welcome to Samba! 🎵\n\n"
		s += "To get started, we need to authenticate with Spotify.\n\n"
		s += faint.Render("Press Enter or Space to begin authentication")
		s += "\n\n" + faint.Render("q - quit")
		return s

	case authURLView:
		s += "🔐 Spotify Authentication\n\n"
		s += wrapText("Please open this URL in your browser to authorize the application:", width) + "\n\n"

		// Make the URL clickable and wrap it
		wrappedURL := wrapText(m.authURL, width)
		s += linkStyle.Render(wrappedURL) + "\n\n"

		s += wrapText("The callback server is running on http://127.0.0.1:8080/callback", width) + "\n"
		s += wrapText("After authorization, you'll be redirected automatically.", width) + "\n\n"
		s += faint.Render("r - regenerate URL • q - quit")
		return s

	case authCodeView:
		s += "🔐 Manual Authorization Code Entry\n\n"
		s += "If the automatic callback didn't work, please enter the authorization code manually:\n\n"
		s += m.textinput.View() + "\n\n"
		s += faint.Render("Enter - submit code • Esc - back • q - quit")
		return s

	case authLoadingView:
		s += "🔄 Authenticating...\n\n"
		s += "Please wait while we exchange your authorization code for an access token.\n\n"
		s += faint.Render("q - quit")
		return s

	case bodyView:
		s += "Queue Info:\n\n"
		s += m.textarea.View() + "\n\n"
		s += faint.Render("esc - back to list")
		return s

	case nowPlayingView:
		if m.user != nil {
			s += "Welcome, " + m.user.GetDisplayName() + "!\n\n"
		}

		// Display currently playing track
		if m.currentlyPlaying != nil && m.currentlyPlaying.Item != nil {
			var trackName, artistName string
			if m.currentlyPlaying.Item.TrackObject != nil {
				trackName = m.currentlyPlaying.Item.TrackObject.GetName()
				if len(m.currentlyPlaying.Item.TrackObject.Artists) > 0 {
					artistName = m.currentlyPlaying.Item.TrackObject.Artists[0].GetName()
				}
			} else if m.currentlyPlaying.Item.EpisodeObject != nil {
				trackName = m.currentlyPlaying.Item.EpisodeObject.GetName()
				artistName = m.currentlyPlaying.Item.EpisodeObject.Show.GetName()
			}

			if trackName != "" {
				isPlaying := m.currentlyPlaying.GetIsPlaying()
				playingStatus := "⏸️ PAUSED"
				musicNote := "🎵"
				if isPlaying {
					playingStatus = "▶️ PLAYING"
					musicNote = "🎶"
				}

				// Create the now playing content
				nowPlayingContent := fmt.Sprintf("%s NOW %s\n\n%s %s\n%s %s", 
					musicNote, playingStatus,
					"♪", nowPlayingTitleStyle.Render(trackName),
					"♫", nowPlayingArtistStyle.Render(artistName))
				
				nowPlayingBox := nowPlayingBoxStyle.Render(nowPlayingContent)
				s += nowPlayingBox + "\n"
			} else {
				s += faint.Render("No track currently playing") + "\n\n"
			}
		} else {
			s += faint.Render("No track currently playing") + "\n\n"
		}

		s += faint.Render("q - view queue • r - refresh • ctrl+c - quit")
		return s

	case listView:
		if m.user != nil {
			s += "Welcome, " + m.user.GetDisplayName() + "!\n\n"
		}

		// Display currently playing track
		if m.currentlyPlaying != nil && m.currentlyPlaying.Item != nil {
			var trackName, artistName string
			if m.currentlyPlaying.Item.TrackObject != nil {
				trackName = m.currentlyPlaying.Item.TrackObject.GetName()
				if len(m.currentlyPlaying.Item.TrackObject.Artists) > 0 {
					artistName = m.currentlyPlaying.Item.TrackObject.Artists[0].GetName()
				}
			} else if m.currentlyPlaying.Item.EpisodeObject != nil {
				trackName = m.currentlyPlaying.Item.EpisodeObject.GetName()
				artistName = m.currentlyPlaying.Item.EpisodeObject.Show.GetName()
			}

			if trackName != "" {
				isPlaying := m.currentlyPlaying.GetIsPlaying()
				playingStatus := "⏸️ PAUSED"
				musicNote := "🎵"
				if isPlaying {
					playingStatus = "▶️ PLAYING"
					musicNote = "🎶"
				}

				// Create the now playing content
				nowPlayingContent := fmt.Sprintf("%s NOW %s\n\n%s %s\n%s %s", 
					musicNote, playingStatus,
					"♪", nowPlayingTitleStyle.Render(trackName),
					"♫", nowPlayingArtistStyle.Render(artistName))
				
				nowPlayingBox := nowPlayingBoxStyle.Render(nowPlayingContent)
				s += nowPlayingBox + "\n"
			}
		}

		// Queue header
		s += queueHeaderStyle.Render("🎼 UPCOMING IN QUEUE") + "\n"

		queue := m.GetQueue()
		totalItems := len(queue)
		currentPage := m.listIndex / itemsPerPage
		startIdx := currentPage * itemsPerPage
		endIdx := startIdx + itemsPerPage
		if endIdx > totalItems {
			endIdx = totalItems
		}

		for i := startIdx; i < endIdx; i++ {
			item := queue[i]
			prefix := " "
			if i == m.listIndex {
				prefix = ">"
			}

			var itemName, artistName string
			if item.TrackObject != nil {
				itemName = item.TrackObject.GetName()
				if len(item.TrackObject.Artists) > 0 {
					artistName = item.TrackObject.Artists[0].GetName()
				}
			} else if item.EpisodeObject != nil {
				itemName = item.EpisodeObject.GetName()
				artistName = item.EpisodeObject.Show.GetName()
			}

			// Wrap long track names and artist names
			trackDisplay := itemName + " | " + artistName
			wrappedTrack := wrapText(trackDisplay, width-2) // account for prefix
			lines := strings.Split(wrappedTrack, "\n")

			// First line gets the prefix, subsequent lines are indented (with faded styling)
			if i == m.listIndex {
				s += queueEnumeratorSelectedStyle.Render(prefix) + queueItemSelectedStyle.Render(lines[0]) + "\n"
			} else {
				s += queueEnumeratorStyle.Render(prefix) + queueItemStyle.Render(lines[0]) + "\n"
			}
			for j := 1; j < len(lines); j++ {
				if i == m.listIndex {
					s += strings.Repeat(" ", 2) + queueItemSelectedStyle.Render(lines[j]) + "\n"
				} else {
					s += strings.Repeat(" ", 2) + queueItemStyle.Render(lines[j]) + "\n"
				}
			}
			s += "\n"
		}

		if totalItems > itemsPerPage {
			totalPages := (totalItems + itemsPerPage - 1) / itemsPerPage
			s += faint.Render("Page ") + faint.Render(fmt.Sprintf("%d", currentPage+1)) + faint.Render(" of ") + faint.Render(fmt.Sprintf("%d", totalPages)) + "\n"
		}
		s += faint.Render("↑↓/kj - navigate • ←→/hl - page • b - back to now playing • r - refresh • q - quit")
	}

	return s
}
