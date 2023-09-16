package lifxAPI

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/JoshKoiro/GoLIFX-Light-Teams-Presence/config"
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

func SetLight(configuration *config.Config, color string) error {

	power := "on"
	if color == "off" {
		power = "off"
		color = "green"
	}

	selectorURI := config.GetSelctorURI("config.yaml")

	url := "https://api.lifx.com/v1/lights/" + selectorURI + "/state"

	duration := configuration.LightSettings.ColorChangeSpeed
	fast := false
	brightness := configuration.LightSettings.Brightness

	// Use fmt.Sprintf to format the string with the variables
	stringPayload := fmt.Sprintf("{\"duration\":%f,\"fast\":%v,\"power\":\"%s\",\"color\":\"%s\",\"brightness\":%f}",
		duration, fast, power, color, brightness)

	payload := strings.NewReader(stringPayload)

	req, err := http.NewRequest("PUT", url, payload)
	if err != nil {
		return fmt.Errorf("error creating LIFX API request: %w", err)
	}

	req.Header.Add("accept", "text/plain")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Authorization", "Bearer "+configuration.LifxAPI.Key)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("error making LIFX API request: %w", err)
	}

	defer res.Body.Close()
	return nil
}
