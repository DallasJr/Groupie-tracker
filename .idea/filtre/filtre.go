package filtre

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Define your filter struct to hold filter values
type Filters struct {
	CreationDateMin float64
	CreationDateMax float64
	FirstAlbumMin   float64
	FirstAlbumMax   float64
	NumMembers      int
	Locations       []string
}

// Function to create a new Filters instance
func NewFilters() *Filters {
	return &Filters{}
}

// Function to create a UI for filters
func CreateFiltersUI(f *Filters) fyne.CanvasObject {
	creationDateRange := container.NewVBox(
		widget.NewLabel("Creation Date Range"),
		widget.NewSlider(0, 2024),
		widget.NewSlider(0, 2024),
	)

	firstAlbumRange := container.NewVBox(
		widget.NewLabel("First Album Date Range"),
		widget.NewSlider(0, 2024),
		widget.NewSlider(0, 2024),
	)

	numMembers := container.NewVBox(
		widget.NewLabel("Number of Members"),
		widget.NewEntry(),
	)

	locations := container.NewVBox(
		widget.NewLabel("Locations"),
		widget.NewCheck("USA", func(checked bool) { /* Update filter values */ }),
		widget.NewCheck("UK", func(checked bool) { /* Update filter values */ }),
		// Add more locations as needed
	)

	applyButton := widget.NewButton("Apply Filters", func() {
		// Handle applying filters
	})

	resetButton := widget.NewButton("Reset Filters", func() {
		// Handle resetting filters
	})

	return container.NewVBox(
		creationDateRange,
		firstAlbumRange,
		numMembers,
		locations,
		container.NewHBox(applyButton, resetButton),
	)
}
