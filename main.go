package main

import (
	"fmt"
	"time"

	"github.com/JoshKoiro/GoLIFX-Light-Teams-Presence/config"
	"github.com/JoshKoiro/GoLIFX-Light-Teams-Presence/lifxAPI"
	"github.com/JoshKoiro/GoLIFX-Light-Teams-Presence/teamsAPI"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

func getTimestamp() {
	t := time.Now()
	fmt.Printf("time: %v\t", t)
}

func updateLight() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file: %s\n", err)
		return
	}

	// Get the teams API status
	status, teamsAPIError := teamsAPI.GetStatus()
	if teamsAPIError != nil {
		fmt.Printf("Error getting status: %s\n", teamsAPIError)
		fmt.Println("You probably need to check your Graph API key")
		return
	}

	// Print the timestamp
	getTimestamp()

	// Read the config file
	configuration, yamlError := config.ReadYAMLFile("config.yaml")
	if yamlError != nil {
		fmt.Printf("Error reading config file: %s\n", yamlError)
		return
	}

	// determine the color to set the light to
	switch status.Availability {
	case "Available":
		fmt.Printf("Status: ")
		color.Green(status.Availability)
		lifxAPI.SetLight(configuration, configuration.StatusColors.AvailableColor)

	case "Away":
		fmt.Printf("Status: ")
		color.Yellow(status.Availability)
		lifxAPI.SetLight(configuration, configuration.StatusColors.AwayColor)

	case "Busy":
		fmt.Printf("Status: ")
		color.Red(status.Availability)
		lifxAPI.SetLight(configuration, configuration.StatusColors.BusyColor)

	case "DoNotDisturb":
		fmt.Printf("Status: ")
		color.Red(status.Availability)
		lifxAPI.SetLight(configuration, configuration.StatusColors.DoNotDisturbColor)
	case "BeRightBack":
		fmt.Printf("Status: ")
		color.Yellow(status.Availability)
		lifxAPI.SetLight(configuration, configuration.StatusColors.BeRightBackColor)
	case "Offline":
		fmt.Printf("Status: ")
		color.White(status.Availability)
		lifxAPI.SetLight(configuration, configuration.StatusColors.OfflineColor)
	default:
		fmt.Printf("Unknown status: %s\n", status.Availability)
		lifxAPI.SetLight(configuration, configuration.StatusColors.OfflineColor)
	}
}

func mainLoop(tick *time.Ticker) {
	for range tick.C {
		updateLight()
	}
}

func main() {
	configRefresh, err := config.ReadYAMLFile("config.yaml")
	if err != nil {
		fmt.Printf("Error reading config file: %s\n", err)
		return
	}

	initialRefreshRate := configRefresh.ApplicationSettings.RefreshRate
	tick := time.NewTicker(time.Duration(initialRefreshRate) * time.Second)

	updateLight()
	mainLoop(tick)
}
