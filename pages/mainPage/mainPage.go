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
	"groupie-tracker/pages/artistPage"
	"groupie-tracker/structs"
	"image/color"
	"time"
)

func LoadPage(myWindow fyne.Window) {

	titleLabel := canvas.NewText("          Groupie Tracker          ", color.White)
	titleLabel.TextSize = 50

	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Search here")

	resultsContainer := container.NewVBox()
	var searchResults []structs.Artist

	performSearch := func() {
		resultsContainer.RemoveAll()
		searchInput := searchEntry.Text
		searchResults = core.Search(searchInput)
		for _, art := range searchResults {
			art := art
			picture := art.GetImage()
			picture.FillMode = canvas.ImageFillContain
			picture.SetMinSize(fyne.NewSize(100, 100))
			fixedName := art.Name
			artistLabel := widget.NewLabel(fixedName)
			infoButton := widget.NewButton("", func() {
				fmt.Println("Loading " + art.Name)
				artistPage.LoadPage(art, myWindow)
			})
			infoButton.Importance = widget.LowImportance
			infoButton.SetIcon(theme.NavigateNextIcon())
			namepicture := container.NewGridWithColumns(2,
				picture,
				artistLabel,
			)
			resultCard := container.New(layout.NewBorderLayout(nil, nil, nil, nil),
				namepicture,
				infoButton)
			for resultCard.MinSize().Width <= 404 {
				artistLabel.Text += " "
				artistLabel.Refresh()
			}
			resultsContainer.Add(resultCard)
		}

	}
	performSearch()
	var latestSearch = time.Time{}

	searchEntry.OnChanged = func(text string) {
		latestSearch = time.Now()
		go func(thisTime time.Time) {
			time.Sleep(1200 * time.Millisecond)
			if latestSearch != thisTime {
				fmt.Println("ignored", text)
				return
			} else {
				fmt.Println("searched", text)
				performSearch()
			}
		}(latestSearch)
	}

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

	topContainer := container.NewVBox(
		container.NewCenter(titleLabel),
		searchEntry,
	)
	bottomContainer := container.NewHBox(
		filterContainer,
		container.NewVScroll(resultsContainer),
	)
	content := container.NewCenter(container.NewVBox(
		topContainer,
		bottomContainer,
	))

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(800, 500))
}
