package config

import (
	"fmt"
	"os"
	"sync"
	// "os/exec" // Import the os/exec package

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

var (
	config     map[string]string
	yamlConfig YAMLConfig // Struct to hold YAML config
	configOnce sync.Once
)

// YAMLConfig defines the structure of your config.yaml file
type YAMLConfig struct {
	PagerDuty struct {
		ServiceIDs []string `yaml:"service_ids"`
	} `yaml:"pagerduty"`
}

// LoadConfig loads configuration from both .env and config.yaml files.
// .env takes precedence for overlapping keys.
func LoadConfig() error {
	var err error
	configOnce.Do(func() { // Ensure Load is only executed once
		// Load .env file first
		err = godotenv.Load()
		if err != nil && !os.IsNotExist(err) { // Ignore if .env doesn't exist, but log other errors
			fmt.Println("Error loading .env file from config folder:", err)
			// Don't return error here, as YAML config might be used instead or .env might not be required
		}

		config = make(map[string]string)
		if os.Getenv("PAGERDUTY_API_KEY") != "" { // Only load from env if set to avoid overwriting YAML if .env is missing
			config["pagerduty_api_key"] = os.Getenv("PAGERDUTY_API_KEY")
		}


		// Load YAML config
		yamlFile, yamlErr := os.ReadFile("config/config.yaml") // Use os.ReadFile instead of ioutil.ReadFile
		if yamlErr != nil {
			fmt.Println("Error reading config.yaml:", yamlErr)
			return // Return error if YAML file cannot be read
		}

		yamlErr = yaml.Unmarshal(yamlFile, &yamlConfig)
		if yamlErr != nil {
			fmt.Println("Error unmarshalling config.yaml:", yamlErr)
			err = yamlErr // Set the error to return from LoadConfig
			return      // Return error if YAML unmarshalling fails
		}
		if config["pagerduty_api_key"] == "" { // If not set by .env, try to get from YAML if available
			config["pagerduty_api_key"] = os.Getenv("PAGERDUTY_API_KEY") // Fallback to env var if not in YAML (or .env)

		}

	})
	return err // Return error from LoadConfig function (if any YAML error occurred)
}

// GetPagerDutyAPIKey retrieves the PagerDuty API key.
func GetPagerDutyAPIKey() string {
	if config == nil {
		if err := LoadConfig(); err != nil {
			fmt.Println("Configuration not loaded and failed to load:", err)
			return ""
		}
	}
	return config["pagerduty_api_key"]
}

// GetServiceIDsFromYAML retrieves the Service IDs from the YAML config.
func GetServiceIDsFromYAML() []string {
	if yamlConfig.PagerDuty.ServiceIDs == nil {
		if err := LoadConfig(); err != nil { // Attempt to load config if not loaded yet
			fmt.Println("Configuration not loaded and failed to load:", err)
			return []string{""} // Return default empty slice in case of error
		}
	}
	if yamlConfig.PagerDuty.ServiceIDs == nil { // Double check after load in case of YAML load failure
		return []string{""} // Return default if still nil after load attempt
	}
	return yamlConfig.PagerDuty.ServiceIDs
}