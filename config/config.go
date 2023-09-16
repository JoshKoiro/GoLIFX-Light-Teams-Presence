package config

import (
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	ApplicationSettings struct {
		RefreshRate int `yaml:"refresh-rate"`
	} `yaml:"application-settings"`
	LifxAPI struct {
		Key string `yaml:"key"`
	} `yaml:"lifx-api"`
	StatusColors struct {
		AvailableColor    string `yaml:"available-color"`
		BusyColor         string `yaml:"busy-color"`
		DoNotDisturbColor string `yaml:"do-not-disturb-color"`
		AwayColor         string `yaml:"away-color"`
		BeRightBackColor  string `yaml:"be-right-back-color"`
		OfflineColor      string `yaml:"offline-color"`
	} `yaml:"status-colors"`
	LightSettings struct {
		Label            string  `yaml:"label"`
		Id               string  `yaml:"id"`
		Brightness       float64 `yaml:"brightness"`
		ColorChangeSpeed float64 `yaml:"color-change-speed"`
	} `yaml:"light-settings"`
}

func ReadYAMLFile(filepath string) (*Config, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := yaml.Unmarshal(content, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func GetSelctorURI(filepath string) string {
	config, _ := ReadYAMLFile(filepath)

	return "label%3A" + strings.ReplaceAll(config.LightSettings.Label, " ", "%20")
}
