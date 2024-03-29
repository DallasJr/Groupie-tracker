package mainPage

import (
	"fmt"
	"groupie-tracker/core"
	"groupie-tracker/structs"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
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
			picture := art.GetImage()
			picture.FillMode = canvas.ImageFillContain
			picture.SetMinSize(fyne.NewSize(100, 100))
			fixedName := art.Name
			artistLabel := widget.NewLabel(fixedName)
			infoButton := widget.NewButton("", func() {

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

	sliderCreationDateStart := widget.NewSlider(0, 2024)
	sliderCreationDateEnd := widget.NewSlider(0, 2024)

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

	creationDateRange := container.NewVBox(
		widget.NewLabel("Creation Date Range"),
		sliderCreationDateStart,
		sliderCreationDateEnd,
	)

	sliderFirstAlbumStart := widget.NewSlider(0, 2024)
	sliderFirstAlbumEnd := widget.NewSlider(0, 2024)

	firstAlbumRange := container.NewVBox(
		widget.NewLabel("First Album Date Range"),
		sliderFirstAlbumStart,
		sliderFirstAlbumEnd,
	)

	entryNumMembers := widget.NewEntry()

	numMembers := container.NewVBox(
		widget.NewLabel("Number of Members"),
		entryNumMembers,
	)

	checkUSA := widget.NewCheck("USA", func(checked bool) {})
	checkUK := widget.NewCheck("UK", func(checked bool) {})
	checkFR := widget.NewCheck("FR", func(checked bool) {})

	locations := container.NewVBox(
		widget.NewLabel("Locations"),
		container.NewHBox(
			checkUSA,
			checkUK,
			checkFR,
		),
	)

	applyButton := widget.NewButton("Apply Filters", func() {

	})

	resetButton := widget.NewButton("Reset Filters", func() {

		sliderCreationDateStart.SetValue(0)
		sliderCreationDateEnd.SetValue(0)
		sliderFirstAlbumStart.SetValue(0)
		sliderFirstAlbumEnd.SetValue(0)
		entryNumMembers.SetText("")
		checkUSA.SetChecked(false)
		checkUK.SetChecked(false)
		checkFR.SetChecked(false)
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
