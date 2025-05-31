package spotify

import (
    "context"
    "fmt"
    "net/http"
)

// SpotifyClient wraps the generated client with auth
type SpotifyClient struct {
    client *APIClient
    token  string
}

// NewClient creates a new Spotify client with authentication
func NewClient(clientID, clientSecret string) (*SpotifyClient, error) {
    cfg := NewConfiguration()
    client := NewAPIClient(cfg)
    
    // TODO: Implement OAuth flow here
    // This is a placeholder - you'll need to implement proper auth
    
    return &SpotifyClient{
        client: client,
    }, nil
}

// Example method - you'll expand this based on your needs
func (s *SpotifyClient) GetQueue(ctx context.Context) (*QueueResponse, error) {
    // Add authorization header
    auth := context.WithValue(ctx, ContextAccessToken, s.token)
    
    queue, resp, err := s.client.PlayerAPI.GetQueue(auth).Execute()
    if err != nil {
        return nil, fmt.Errorf("failed to get queue: %w", err)
    }
    defer resp.Body.Close()
    
    return queue, nil
}
