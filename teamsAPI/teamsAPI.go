package teamsAPI

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type PresenceResponse struct {
	Availability string `json:"availability"`
	Activity     string `json:"activity"`
}

func GetStatus() (PresenceResponse, error) {
	// Import values from .env file
	err := godotenv.Load()
	if err != nil {
		return PresenceResponse{}, fmt.Errorf("error loading .env file: %w", err)
	}

	// Get authorization key from .env file
	key := os.Getenv("GRAPH_API_KEY")

	// Create request
	req, err := http.NewRequest("GET", "https://graph.microsoft.com/beta/me/presence", nil)
	if err != nil {
		return PresenceResponse{}, fmt.Errorf("error creating request: %w", err)
	}

	// Add headers to request
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+key)

	// Send request and get response
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return PresenceResponse{}, fmt.Errorf("error sending request: %w", err)
	}

	// collect response body
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return PresenceResponse{}, fmt.Errorf("error reading response body: %w", err)
	}

	// parse response body
	availability, activity, err := getPresence(string(body))
	if err != nil {
		return PresenceResponse{}, fmt.Errorf("error parsing response body: %w", err)
	}

	// Return response body
	return PresenceResponse{Availability: availability, Activity: activity}, err
}

func getPresence(responseJSON string) (string, string, error) {
	var presence PresenceResponse

	err := json.Unmarshal([]byte(responseJSON), &presence)
	if err != nil {
		return "", "", fmt.Errorf("error parsing response body: %w", err)
	}

	return presence.Availability, presence.Activity, err
}
