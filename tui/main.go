package main

import (
	"context"
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/oauth2"

	spotify "github.com/ericflores108/samba/pkg/spotify"
)

func main() {
	// Setup OAuth2 configuration
	config := &oauth2.Config{
		ClientID:     os.Getenv("SPOTIFY_CLIENT_ID"),
		ClientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
		RedirectURL:  "http://localhost:8080/callback",
		Scopes:       []string{"user-read-private", "user-read-email", "user-library-read"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.spotify.com/authorize",
			TokenURL: "https://accounts.spotify.com/api/token",
		},
	}

	if config.ClientID == "" || config.ClientSecret == "" {
		log.Fatal("Please set SPOTIFY_CLIENT_ID and SPOTIFY_CLIENT_SECRET environment variables")
	}

	// Get access token (this is a simplified version - in a real app you'd handle the OAuth flow properly)
	token := &oauth2.Token{
		AccessToken: os.Getenv("SPOTIFY_ACCESS_TOKEN"),
		TokenType:   "Bearer",
	}

	if token.AccessToken == "" {
		fmt.Println("Please set SPOTIFY_ACCESS_TOKEN environment variable")
		fmt.Println("You can get one from: https://developer.spotify.com/console/get-current-user/")
		os.Exit(1)
	}

	// Create Spotify API client
	ctx := context.WithValue(context.Background(), spotify.ContextOAuth2, oauth2.StaticTokenSource(token))
	cfg := spotify.NewConfiguration()
	client := spotify.NewAPIClient(cfg)

	// Get current user
	user, _, err := client.UsersAPI.GetCurrentUsersProfile(ctx).Execute()
	if err != nil {
		log.Fatalf("Failed to get current user: %v", err)
	}

	m := NewModel(user, client, ctx)

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		log.Fatalf("p.Run: %v", err)
	}
}
