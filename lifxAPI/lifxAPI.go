package lifxAPI

import (
	"fmt"
	"io"
	"net/http"
)

func GetLights(token string) ([]byte, error) {
	url := "https://api.lifx.com/v1/lights/all"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating LIFX API request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making LIFX API request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body from LIFX API: %w", err)
	}

	return body, nil
}
