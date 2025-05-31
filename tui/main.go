package main

import (
	"context"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"

	spotify "github.com/ericflores108/samba/pkg/spotifywrapper"
)

func main() {
	// Load environment variables from .env file if not in production
	if os.Getenv("GO_ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Println("No .env file found, using system environment variables")
		}
	}

	clientID := os.Getenv("SPOTIFY_CLIENT_ID")
	clientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET")
	redirectURL := "http://127.0.0.1:8080/callback"
	scopes := []string{"user-read-playback-state", "user-read-currently-playing", "user-modify-playback-state", "user-read-private", "user-read-email"}

	if clientID == "" || clientSecret == "" {
		log.Fatal("Please set SPOTIFY_CLIENT_ID and SPOTIFY_CLIENT_SECRET environment variables")
	}

	// Create Spotify client with OAuth wrapper
	spotifyClient, err := spotify.NewClient(clientID, clientSecret, redirectURL, scopes)
	if err != nil {
		log.Fatalf("Failed to create Spotify client: %v", err)
	}

	ctx := context.Background()

	// Check if we have a stored access token
	accessToken := os.Getenv("SPOTIFY_ACCESS_TOKEN")
	refreshToken := os.Getenv("SPOTIFY_REFRESH_TOKEN")
	if accessToken != "" {
		// Use existing token
		token := &oauth2.Token{
			AccessToken:  accessToken,
			TokenType:    "Bearer",
			RefreshToken: refreshToken,
		}
		spotifyClient.SetToken(token)

		// Verify token works by getting user profile
		user, _, err := spotifyClient.GetCurrentUserProfile(ctx)
		if err != nil {
			log.Printf("Failed to get user profile: %v", err)
			// Token might be invalid, start interactive auth
			m := NewModel(spotifyClient, ctx)
			p := tea.NewProgram(m)
			if _, err := p.Run(); err != nil {
				log.Fatalf("p.Run: %v", err)
			}
			return
		}

		// Token works, go directly to main app
		m := NewAuthenticatedModel(user, spotifyClient, ctx)
		p := tea.NewProgram(m)
		if _, err := p.Run(); err != nil {
			log.Fatalf("p.Run: %v", err)
		}
		return
	}

	// No token available, start interactive auth flow
	m := NewModel(spotifyClient, ctx)
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		log.Fatalf("p.Run: %v", err)
	}
}
