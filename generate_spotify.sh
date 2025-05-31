#!/bin/bash

# generate_spotify.sh - Generate Spotify OpenAPI client

set -e

# Configuration
SPOTIFY_OPENAPI_URL="https://developer.spotify.com/reference/web-api/open-api-schema.yaml"
OUTPUT_DIR="./pkg/spotify"
PACKAGE_NAME="spotify"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${GREEN}🎵 Generating Spotify OpenAPI client...${NC}"

# Check if openapi-generator is installed
if ! command -v openapi-generator &> /dev/null; then
    echo -e "${YELLOW}Installing openapi-generator...${NC}"
    
    # Try different installation methods
    if command -v brew &> /dev/null; then
        brew install openapi-generator
    elif command -v npm &> /dev/null; then
        npm install -g @openapitools/openapi-generator-cli
        # Create alias for consistency
        alias openapi-generator=openapi-generator-cli
    else
        echo -e "${RED}Error: Please install openapi-generator${NC}"
        echo "Visit: https://openapi-generator.tech/docs/installation"
        echo "Homebrew: brew install openapi-generator"
        echo "NPM: npm install -g @openapitools/openapi-generator-cli"
        exit 1
    fi
fi

# Create output directory
mkdir -p "$OUTPUT_DIR"

# Generate the client
echo -e "${GREEN}Generating Go client...${NC}"
openapi-generator generate \
  -i "$SPOTIFY_OPENAPI_URL" \
  -g go \
  -o "$OUTPUT_DIR" \
  --package-name "$PACKAGE_NAME" \
  --git-user-id "ericflores108" \
  --git-repo-id "samba" \
  --additional-properties=enumClassPrefix=true,structPrefix=true,generateInterfaces=true \
  --skip-validate-spec

# Clean up unnecessary files
echo -e "${GREEN}Cleaning up...${NC}"
rm -rf "$OUTPUT_DIR"/.openapi-generator
rm -f "$OUTPUT_DIR"/.openapi-generator-ignore
rm -f "$OUTPUT_DIR"/go.mod "$OUTPUT_DIR"/go.sum
rm -rf "$OUTPUT_DIR"/test

# Create a simple wrapper example
cat > "$OUTPUT_DIR"/wrapper.go << 'EOF'
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
EOF

echo -e "${GREEN}✅ Spotify API client generated successfully!${NC}"
echo -e "${YELLOW}Next steps:${NC}"
echo "1. Implement OAuth authentication in wrapper.go"
echo "2. Add specific methods you need for your queue operations"
echo "3. Import in your server and TUI: import \"github.com/ericflores108/samba/pkg/spotify\""

# Make the script executable
chmod +x "$0"