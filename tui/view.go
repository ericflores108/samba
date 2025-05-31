package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	appNameStyle = lipgloss.NewStyle().Background(lipgloss.Color("99")).Padding(0, 1)

	faint = lipgloss.NewStyle().Foreground(lipgloss.Color("255")).Faint(true)

	listEnumeratorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("99")).MarginRight(1)
)

func (m model) View() string {
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
		s += "Please open this URL in your browser to authorize the application:\n\n"
		s += lipgloss.NewStyle().Foreground(lipgloss.Color("12")).Render(m.authURL) + "\n\n"
		s += "The callback server is running on http://127.0.0.1:8080/callback\n"
		s += "After authorization, you'll be redirected automatically.\n\n"
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

	case listView:
		if m.user != nil {
			s += "Welcome, " + m.user.GetDisplayName() + "!\n\n"
		}

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

			s += listEnumeratorStyle.Render(prefix) + itemName + " | " + faint.Render(artistName) + "\n\n"
		}

		if totalItems > itemsPerPage {
			totalPages := (totalItems + itemsPerPage - 1) / itemsPerPage
			s += faint.Render("Page ") + faint.Render(fmt.Sprintf("%d", currentPage+1)) + faint.Render(" of ") + faint.Render(fmt.Sprintf("%d", totalPages)) + "\n"
		}
		s += faint.Render("↑↓/kj - navigate • ←→/hl - page • r - refresh • q - quit")
	}

	return s
}
