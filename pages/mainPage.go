package pages

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

func LoadMainPage(myWindow fyne.Window) {

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
			artistLabel := widget.NewLabel(art.Name)
			button := widget.NewButton("", func() {
				LoadArtistPage(art, myWindow)
			})
			button.Importance = widget.LowImportance
			button.SetIcon(theme.NavigateNextIcon())
			namepicture := container.NewGridWithColumns(2,
				picture,
				artistLabel,
			)
			resultCard := container.New(layout.NewBorderLayout(nil, nil, nil, nil),
				namepicture,
				button)
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
	// Déclaration et initialisation des sliders pour Creation Date Range
	sliderCreationDateStart := widget.NewSlider(0, 2024)
	sliderCreationDateEnd := widget.NewSlider(0, 2024)

	creationDateRange := container.NewVBox(
		widget.NewLabel("Creation Date Range"),
		sliderCreationDateStart,
		sliderCreationDateEnd,
	)

	// Déclaration et initialisation des sliders pour First Album Date Range
	sliderFirstAlbumStart := widget.NewSlider(0, 2024)
	sliderFirstAlbumEnd := widget.NewSlider(0, 2024)

	firstAlbumRange := container.NewVBox(
		widget.NewLabel("First Album Date Range"),
		sliderFirstAlbumStart,
		sliderFirstAlbumEnd,
	)

	// Déclaration et initialisation de l'entry pour Number of Members
	entryNumMembers := widget.NewEntry()

	numMembers := container.NewVBox(
		widget.NewLabel("Number of Members"),
		entryNumMembers,
	)

	// Déclaration et initialisation des checkboxes pour Locations
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
		// Ici, vous pouvez maintenant accéder directement aux variables
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

	favoriteContainer := container.NewVBox()
	favoritesLabel := canvas.NewText("Favorites:", color.White)
	favoritesLabel.TextSize = 30
	favoriteContainer.Add(container.NewCenter(favoritesLabel))
	if len(core.Favorites) > 0 {
		favorites := container.NewGridWithColumns(4)
		for _, favorite := range core.Favorites {
			artist := structs.GetArtist(favorite)
			image := artist.GetImage()
			artistLabel := widget.NewLabel(structs.GetArtist(favorite).Name)

			card := widget.NewCard("", "", container.NewVBox(image, container.NewCenter(artistLabel)))

			button := widget.NewButton("", func() {
				LoadArtistPage(artist, myWindow)
			})
			button.Importance = widget.LowImportance
			button.SetIcon(theme.NavigateNextIcon())
			finalCard := container.New(layout.NewBorderLayout(nil, nil, nil, nil),
				card,
				container.NewGridWithColumns(3, layout.NewSpacer(), button, layout.NewSpacer()))
			favorites.Add(finalCard)
		}
		favoriteContainer.Add(favorites)
	}
	final := container.NewVBox(content, favoriteContainer)
	myWindow.SetContent(container.NewVScroll(final))
}
