package mainPage

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"groupie-tracker/core"
	"groupie-tracker/structs"
	"image/color"
)

func LoadPage(myWindow fyne.Window) {

	titleLabel := canvas.NewText("Groupie Tracker", color.White)
	titleLabel.TextSize = 50
	titleContainer := container.NewCenter(titleLabel)

	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Search here")

	resultsContainer := container.NewVBox()
	var searchResults []structs.Artist
	searchButton := widget.NewButtonWithIcon("", theme.MailForwardIcon(), func() {
		resultsContainer.RemoveAll()
		searchInput := searchEntry.Text
		searchResults = core.SearchArtistsByName(searchInput)
		for _, art := range searchResults {

			picture := art.GetImage()
			picture.FillMode = canvas.ImageFillContain
			picture.SetMinSize(fyne.NewSize(100, 100))
			resultCard := container.NewHBox(
				container.New(layout.NewCenterLayout(), picture),
				container.NewWithoutLayout(widget.NewLabel(art.Name)),
			)
			/*resultCard.SetOnTapped(func() {
				// Handle the tap event here
				label.SetText("Button Clicked!")
			})*/
			resultsContainer.Add(resultCard)
			fmt.Println("Found: " + art.Name)
		}
	})

	searchEntry.Resize(fyne.NewSize(415, searchEntry.MinSize().Height))
	searchEntry.Move(fyne.NewPos(-40, 0))
	searchButton.Resize(fyne.NewSize(searchButton.MinSize().Width, searchButton.MinSize().Height))
	searchButton.Move(fyne.NewPos(380, 0))

	inputContainer := container.NewWithoutLayout(searchEntry, searchButton)

	//Filters:
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
	//resultsContainer.Resize(fyne.NewSize(415, resultsContainer.MinSize().Height))

	//searchEntry.Resize(fyne.NewSize(415, searchEntry.MinSize().Height))

	bottomContainer := container.NewHBox(
		filterContainer,
		resultsContainer,
	)

	content := container.NewCenter(container.NewVBox(
		titleContainer,
		inputContainer,
		bottomContainer,
	))

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(800, 500))
}
