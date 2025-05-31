package spotify

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"golang.org/x/oauth2"
)

// OAuthConfig holds OAuth configuration
type OAuthConfig struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
	Scopes       []string
}

// TokenInfo holds token information
type TokenInfo struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

// SpotifyClient wraps the generated client with auth
type SpotifyClient struct {
	client      *APIClient
	oauthConfig *OAuthConfig
	token       *oauth2.Token
}

// NewClient creates a new Spotify client with OAuth configuration
func NewClient(clientID, clientSecret, redirectURL string, scopes []string) (*SpotifyClient, error) {
	cfg := NewConfiguration()
	client := NewAPIClient(cfg)

	oauthConfig := &OAuthConfig{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes:       scopes,
	}

	return &SpotifyClient{
		client:      client,
		oauthConfig: oauthConfig,
	}, nil
}

// GetAuthURL generates the OAuth authorization URL
func (sc *SpotifyClient) GetAuthURL() (string, string, error) {
	state := generateRandomString(16)

	params := url.Values{}
	params.Add("response_type", "code")
	params.Add("client_id", sc.oauthConfig.ClientID)
	params.Add("scope", strings.Join(sc.oauthConfig.Scopes, " "))
	params.Add("redirect_uri", sc.oauthConfig.RedirectURL)
	params.Add("state", state)

	authURL := "https://accounts.spotify.com/authorize?" + params.Encode()
	return authURL, state, nil
}

// ExchangeCodeForToken exchanges authorization code for access token
func (sc *SpotifyClient) ExchangeCodeForToken(code, state string) error {
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("redirect_uri", sc.oauthConfig.RedirectURL)

	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))
	if err != nil {
		return fmt.Errorf("failed to create token request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(sc.oauthConfig.ClientID+":"+sc.oauthConfig.ClientSecret)))

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to exchange code for token: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("token exchange failed with status: %d", resp.StatusCode)
	}

	var tokenInfo TokenInfo
	if err := json.NewDecoder(resp.Body).Decode(&tokenInfo); err != nil {
		return fmt.Errorf("failed to decode token response: %w", err)
	}

	// Convert to oauth2.Token
	sc.token = &oauth2.Token{
		AccessToken:  tokenInfo.AccessToken,
		TokenType:    tokenInfo.TokenType,
		RefreshToken: tokenInfo.RefreshToken,
		Expiry:       time.Now().Add(time.Duration(tokenInfo.ExpiresIn) * time.Second),
	}

	return nil
}

// RefreshToken refreshes the access token using the refresh token
func (sc *SpotifyClient) RefreshToken() error {
	if sc.token == nil || sc.token.RefreshToken == "" {
		return fmt.Errorf("no refresh token available")
	}

	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", sc.token.RefreshToken)

	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))
	if err != nil {
		return fmt.Errorf("failed to create refresh request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(sc.oauthConfig.ClientID+":"+sc.oauthConfig.ClientSecret)))

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to refresh token: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("token refresh failed with status: %d", resp.StatusCode)
	}

	var tokenInfo TokenInfo
	if err := json.NewDecoder(resp.Body).Decode(&tokenInfo); err != nil {
		return fmt.Errorf("failed to decode refresh response: %w", err)
	}

	// Update token
	sc.token.AccessToken = tokenInfo.AccessToken
	sc.token.Expiry = time.Now().Add(time.Duration(tokenInfo.ExpiresIn) * time.Second)

	// Refresh token might be returned, update if present
	if tokenInfo.RefreshToken != "" {
		sc.token.RefreshToken = tokenInfo.RefreshToken
	}

	return nil
}

// IsTokenValid checks if the current token is valid and not expired
func (sc *SpotifyClient) IsTokenValid() bool {
	return sc.token != nil && sc.token.Valid()
}

// ensureValidToken ensures we have a valid token, refreshing if necessary
func (sc *SpotifyClient) ensureValidToken() error {
	if sc.token == nil {
		return fmt.Errorf("no token available, please authenticate first")
	}

	if !sc.token.Valid() {
		return sc.RefreshToken()
	}

	return nil
}

// GetAuthenticatedContext returns a context with the OAuth token for API calls
func (sc *SpotifyClient) GetAuthenticatedContext(ctx context.Context) (context.Context, error) {
	if err := sc.ensureValidToken(); err != nil {
		return nil, err
	}

	// Create a TokenSource from the token
	tokenSource := oauth2.StaticTokenSource(sc.token)
	return context.WithValue(ctx, ContextOAuth2, tokenSource), nil
}

// SetToken allows setting a token directly (useful for stored tokens)
func (sc *SpotifyClient) SetToken(token *oauth2.Token) {
	sc.token = token
}

// GetToken returns the current token (useful for storage)
func (sc *SpotifyClient) GetToken() *oauth2.Token {
	return sc.token
}

// GetAPIClient returns the underlying API client
func (sc *SpotifyClient) GetAPIClient() *APIClient {
	return sc.client
}

// Helper methods for common API operations with automatic authentication

// GetCurrentUserProfile gets the current user's profile
func (sc *SpotifyClient) GetCurrentUserProfile(ctx context.Context) (*PrivateUserObject, *http.Response, error) {
	authCtx, err := sc.GetAuthenticatedContext(ctx)
	if err != nil {
		return nil, nil, err
	}

	result, resp, err := sc.client.UsersAPI.GetCurrentUsersProfile(authCtx).Execute()
	return result, resp, err
}

// SearchItems searches for tracks, artists, albums, etc.
func (sc *SpotifyClient) SearchItems(ctx context.Context, q string, types []string, limit *int32, offset *int32) (*Search200Response, *http.Response, error) {
	authCtx, err := sc.GetAuthenticatedContext(ctx)
	if err != nil {
		return nil, nil, err
	}
	req := sc.client.SearchAPI.Search(authCtx).Q(q).Type_(types)
	if limit != nil {
		req = req.Limit(*limit)
	}
	if offset != nil {
		req = req.Offset(*offset)
	}

	result, resp, err := req.Execute()
	return result, resp, err
}

// GetQueue gets the current user's playback queue
func (sc *SpotifyClient) GetQueue(ctx context.Context) (*QueueObject, *http.Response, error) {
	authCtx, err := sc.GetAuthenticatedContext(ctx)
	if err != nil {
		return nil, nil, err
	}

	result, resp, err := sc.client.PlayerAPI.GetQueue(authCtx).Execute()
	return result, resp, err
}

// generateRandomString creates a random string for state parameter
func generateRandomString(length int) string {
	bytes := make([]byte, length)
	rand.Read(bytes)
	return base64.URLEncoding.EncodeToString(bytes)[:length]
}
