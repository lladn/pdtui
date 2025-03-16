package pagerdutyapi

import (
	"context"
	"encoding/json"
	"fmt"
	"pdtui/config"

	"github.com/PagerDuty/go-pagerduty"
)

var newClient = pagerduty.NewClient

func ListIncidents() ([]pagerduty.Incident, error) {
	// Load configuration
	if err := config.LoadConfig("config.yaml"); err != nil {
		return nil, fmt.Errorf("error loading configuration: %w", err)
	}

	apiKey := config.GetPagerDutyAPIKey()
	if apiKey == "" {
		return nil, fmt.Errorf("PagerDuty API key not found in configuration")
	}

	client := newClient(apiKey)
	serviceIDs := config.GetServiceIDs()

	opts := pagerduty.ListIncidentsOptions{
		Statuses:   []string{"triggered", "acknowledged", "resolved"},
		ServiceIDs: serviceIDs,
		Limit:      25,
	}

	incidents, err := client.ListIncidentsWithContext(context.Background(), opts)
	if err != nil {
		return nil, fmt.Errorf("error listing incidents: %w", err)
	}

	jsonData, err := json.MarshalIndent(incidents, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("error marshalling incidents to JSON: %w", err)
	}
	fmt.Println(string(jsonData))

	return incidents.Incidents, nil
}