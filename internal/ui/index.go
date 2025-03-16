package ui

import (
	"fmt"
	"github.com/PagerDuty/go-pagerduty" // Import the PagerDuty package
	"github.com/rivo/tview"
	"pdtui/api"
)

// StartUI initializes and runs the TUI application
func StartUI() {
	app := tview.NewApplication()

	// Fetch incidents from PagerDuty
	incidents, err := pagerdutyapi.ListIncidents()
	if err != nil {
		panic(fmt.Errorf("failed to fetch incidents: %w", err))
	}

	// Create header
	header := tview.NewBox().
		SetBorder(true).
		SetTitle(" PDTUI - PagerDuty Incidents ").
		SetTitleAlign(tview.AlignLeft).
		SetBorderPadding(0, 0, 0, 0)

	// Create left panel with widgets
	leftPanel := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tview.NewBox().
			SetBorder(true).
			SetTitle(" NAVIGATION ").
			SetTitleAlign(tview.AlignLeft).
			SetBorderPadding(0, 0, 0, 0),
			0, 1, false).
		AddItem(tview.NewBox().
			SetBorder(true).
			SetTitle(" SETTINGS ").
			SetTitleAlign(tview.AlignLeft).
			SetBorderPadding(0, 0, 0, 0),
			0, 1, false).
		AddItem(tview.NewBox().
			SetBorder(true).
			SetTitle(" STATUS ").
			SetTitleAlign(tview.AlignLeft).
			SetBorderPadding(0, 0, 0, 0),
			0, 1, false)

	// Create main content area with incidents
	incidentsList := tview.NewTextView().
		SetDynamicColors(true).
		SetScrollable(true).
		SetText(formatIncidents(incidents))

	mainContent := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(incidentsList, 0, 1, false).
		AddItem(tview.NewBox().
			SetBorder(true).
			SetTitle(" DETAILS ").
			SetBorderPadding(0, 0, 0, 0),
			0, 1, false)

	// Main layout
	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(header, 3, 0, false).
		AddItem(tview.NewFlex().
			AddItem(leftPanel, 0, 1, false).
			AddItem(mainContent, 0, 2, false),
			0, 1, false)

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}

// formatIncidents formats the incidents for display
func formatIncidents(incidents []pagerduty.Incident) string {
	result := ""
	for _, incident := range incidents {
		result += fmt.Sprintf("[red]%s[-]\n  Status: %s\n  Service: %s\n\n",
			incident.Title,
			incident.Status,
			incident.Service.Summary,
		)
	}
	return result
}