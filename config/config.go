package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	PagerDuty struct {
		ServiceIDs []string `yaml:"serviceIDs"`
	} `yaml:"pagerDuty"`

	UI struct {
		Header struct {
			Title  string `yaml:"title"`
			Height int    `yaml:"height"`
		} `yaml:"header"`

		LeftPanel struct {
			Widgets []string `yaml:"widgets"`
		} `yaml:"leftPanel"`

		MainContent struct {
			Sections []string `yaml:"sections"`
		} `yaml:"mainContent"`
	} `yaml:"ui"`
}

var cfg Config

// LoadConfig loads the configuration from a file and environment variables
func LoadConfig(path string) error {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("failed to load .env file: %w", err)
	}

	viper.SetConfigFile(path)
	viper.AutomaticEnv() // Override config values with environment variables

	// Set default values
	viper.SetDefault("pagerDuty.serviceIDs", []string{})
	viper.SetDefault("ui.header.title", "PDTUI")
	viper.SetDefault("ui.header.height", 3)
	viper.SetDefault("ui.leftPanel.widgets", []string{"NAVIGATION", "SETTINGS", "STATUS"})
	viper.SetDefault("ui.mainContent.sections", []string{"INCIDENTS", "DETAILS"})

	// Read the config file
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	// Unmarshal the config into a struct
	if err := viper.Unmarshal(&cfg); err != nil {
		return fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return nil
}

// GetPagerDutyAPIKey returns the PagerDuty API key from the environment
func GetPagerDutyAPIKey() string {
	return os.Getenv("PAGERDUTY_API_KEY")
}

// GetServiceIDs returns the PagerDuty service IDs from the configuration
func GetServiceIDs() []string {
	return cfg.PagerDuty.ServiceIDs
}