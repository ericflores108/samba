package main

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	appNameStyle = lipgloss.NewStyle().Background(lipgloss.Color("99")).Padding(0, 1)

	faint = lipgloss.NewStyle().Foreground(lipgloss.Color("255")).Faint(true)

	listEnumeratorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("99")).MarginRight(1)
)

func (m model) View() string {
	s := appNameStyle.Render("SPOTIFY TUI") + "\n\n"

	// No title view for Spotify

	if m.state == bodyView {
		s += "Track Info:\n\n"
		s += m.textarea.View() + "\n\n"
		s += faint.Render("esc - back to list")
	}

	if m.state == listView {
		if m.user != nil {
			s += "Welcome, " + m.user.GetDisplayName() + "!\n\n"
		}
		for i, track := range m.tracks {
			prefix := " "
			if i == m.listIndex {
				prefix = ">"
			}
			trackName := track.Track.GetName()
			artistName := ""
			if len(track.Track.Artists) > 0 {
				artistName = track.Track.Artists[0].GetName()
			}
			s += listEnumeratorStyle.Render(prefix) + trackName + " | " + faint.Render(artistName) + "\n\n"
		}
		s += faint.Render("r - refresh • enter - details • q - quit")
	}

	return s
}
