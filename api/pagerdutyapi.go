package pagerdutyapi

import (
	"context"
	"encoding/json"
	"fmt"
	// "os"

	"github.com/PagerDuty/go-pagerduty"
	"pdtui/config"
)

var newClient = pagerduty.NewClient

func ListIncidents() ([]pagerduty.Incident, error) {
	err := config.LoadConfig() 
	if err != nil {
		fmt.Println("Error loading configuration:", err)
		return nil, fmt.Errorf("configuration error: %w", err)
	}

	apiKey := config.GetPagerDutyAPIKey()
	if apiKey == "" {
		return nil, fmt.Errorf("PagerDuty API key not found in configuration")
	}
	client := newClient(apiKey)
	serviceIDs := config.GetServiceIDsFromYAML() 

	opts := pagerduty.ListIncidentsOptions{
		Statuses:   []string{"acknowledged", "resolved"},
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