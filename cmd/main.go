package main

import (
	"pdtui/internal/ui"
)

func main() {
	// Start the TUI application
	ui.StartUI()
}

// func main() {
// 	incidents, err := pagerdutyapi.ListIncidents()
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
// 		os.Exit(1)
// 	}

// 	fmt.Println("\nIncidents List:")
// 	if len(incidents) == 0 {
// 		fmt.Println("No incidents found with the specified criteria.")
// 	} else {
// 		// Incidents are already printed in JSON format by ListIncidents function
// 		// You can iterate and print them in a different format if needed
// 		// For example, to print incident IDs:
// 		/*
// 		for _, incident := range incidents {
// 			fmt.Println("Incident ID:", incident.ID)
// 		}
// 		*/
// 	}
// }
