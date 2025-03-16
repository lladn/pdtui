package ui

import (
	"github.com/rivo/tview"
)

// StartUI initializes and runs the TUI application
func MonitorUI() {
	app := tview.NewApplication()

	// Create header with title
	header := tview.NewBox().
		SetBorder(true).
		SetTitle(" MY APPLICATION HEADER ")

	// Create left panel with 3 titled widgets
	leftPanel := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tview.NewBox().
			SetBorder(true).
			SetTitle(" NAVIGATION "),
			0, 1, false).
		AddItem(tview.NewBox().
			SetBorder(true).
			SetTitle(" SETTINGS "),
			0, 1, false).
		AddItem(tview.NewBox().
			SetBorder(true).
			SetTitle(" STATUS "),
			0, 1, false)

	// Main content area with 2 vertical boxes
	mainContent := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tview.NewBox().
			SetBorder(true).
			SetTitle(" CONTENT PREVIEW "),
			0, 1, false).
		AddItem(tview.NewBox().
			SetBorder(true).
			SetTitle(" DETAILS "),
			0, 1, false)

	// Main layout
	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(header, 3, 0, false). // Fixed height header
		AddItem(tview.NewFlex().
			AddItem(leftPanel, 0, 1, false).  // Left widgets (1/3 width)
			AddItem(mainContent, 0, 2, false), 0, 1, false) // Main content (2/3 width)

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}