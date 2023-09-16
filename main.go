package main

import (
	"fmt"

	"github.com/JoshKoiro/GoLIFX-Light-Teams-Presence/config"
	"github.com/JoshKoiro/GoLIFX-Light-Teams-Presence/lifxAPI"
	"github.com/JoshKoiro/GoLIFX-Light-Teams-Presence/teamsAPI"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file: %s\n", err)
		return
	}

	// Get the teams API status
	status, err := teamsAPI.GetStatus()
	if err != nil {
		fmt.Printf("Error getting status: %s\n", err)
		return
	}

	// Print the teams API status
	fmt.Printf("Status: %s\n", status.Availability)

	configuration, err := config.ReadYAMLFile("config.yaml")
	if err != nil {
		fmt.Printf("Error reading config file: %s\n", err)
		return
	}

	lighterror := lifxAPI.SetLight(configuration, "purple")
	if lighterror != nil {
		fmt.Printf("Error setting light: %s\n", lighterror)
		return
	}
}
