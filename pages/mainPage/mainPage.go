package mainPage

import (
	"groupie-tracker/pages/artistPage"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func LoadPage(myWindow fyne.Window) {

	titleLabel := canvas.NewText("Groupie Tracker", color.White)
	titleLabel.TextSize = 50
	titleContainer := container.NewCenter(titleLabel)

	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Search here")

	searchButton := widget.NewButtonWithIcon("", theme.MailForwardIcon(), func() {
		artistPage.LoadPage(myWindow)
	})

	searchEntry.Resize(fyne.NewSize(415, searchEntry.MinSize().Height))
	searchEntry.Move(fyne.NewPos(-40, 0))
	searchButton.Resize(fyne.NewSize(searchButton.MinSize().Width, searchButton.MinSize().Height))
	searchButton.Move(fyne.NewPos(380, 0))

	inputContainer := container.NewWithoutLayout(searchEntry, searchButton)

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
		container.NewHBox(
			widget.NewCheck("USA", func(checked bool) {}),
			widget.NewCheck("UK", func(checked bool) {}),
			widget.NewCheck("FR", func(checked bool) {}),
		),
	)

	applyButton := widget.NewButton("Apply Filters", func() {

	})

	resetButton := widget.NewButton("Reset Filters", func() {

	})

	filterContainer := container.NewVBox(
		creationDateRange,
		firstAlbumRange,
		numMembers,
		locations,
		container.NewHBox(applyButton, resetButton),
	)

	content := container.NewCenter(container.NewVBox(
		titleContainer,
		inputContainer,
		filterContainer,
	))
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(800, 500))
}
