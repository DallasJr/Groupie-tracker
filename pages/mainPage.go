package pages

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
	"strings"
)

func LoadMainPage(myWindow fyne.Window) {
	titleLabel := canvas.NewText("          Groupie Tracker          ", color.White)
	titleLabel.TextSize = 50

	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Search here")

	suggestionBox := container.NewVBox()

	updateSuggestions := func(text string) {

		suggestionBox.Objects = nil

		suggestions := core.GetSuggestions(text)
		if len(suggestions) < 10 {
			for _, suggestion := range suggestions {
				suggestion := suggestion
				button := widget.NewButton(suggestion, func() {
					parts := strings.SplitN(suggestion, " - ", 2)
					searchEntry.SetText(parts[0])
				})
				suggestionBox.Add(button)
			}
		}
	}

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
	searchEntry.OnChanged = func(text string) {
		updateSuggestions(text)
		performSearch()
	}

	//Filters:
	// Déclaration et initialisation des sliders pour Creation Date Range
	creationDateSliderMin := widget.NewSlider(float64(core.CreationDateRange[0]), float64(core.CreationDateRange[1]))
	creationDateSliderMax := widget.NewSlider(float64(core.CreationDateRange[0]), float64(core.CreationDateRange[1]))
	creationDateSliderMax.SetValue(float64(core.CreationDateRange[1]))
	firstAlbumSliderMin := widget.NewSlider(float64(core.FirstAlbumDateRange[0]), float64(core.FirstAlbumDateRange[1]))
	firstAlbumSliderMax := widget.NewSlider(float64(core.FirstAlbumDateRange[0]), float64(core.FirstAlbumDateRange[1]))
	firstAlbumSliderMax.SetValue(float64(core.FirstAlbumDateRange[1]))
	numberOfMembersSliderMin := widget.NewSlider(float64(core.NumberOfMembersRange[0]), float64(core.NumberOfMembersRange[1]))
	numberOfMembersSliderMax := widget.NewSlider(float64(core.NumberOfMembersRange[0]), float64(core.NumberOfMembersRange[1]))
	numberOfMembersSliderMax.SetValue(float64(core.NumberOfMembersRange[1]))

	creationDateLabel := widget.NewLabel(fmt.Sprintf("Creation Date Range: %d - %d", core.CreationDateRange[0], core.CreationDateRange[1]))
	firstAlbumLabel := widget.NewLabel(fmt.Sprintf("First Album Date Range: %d - %d", core.FirstAlbumDateRange[0], core.FirstAlbumDateRange[1]))
	numberOfMembersLabel := widget.NewLabel(fmt.Sprintf("Number of Members Range: %d - %d", core.NumberOfMembersRange[0], core.NumberOfMembersRange[1]))

	updateLabels := func() {
		creationDateLabel.SetText(fmt.Sprintf("Creation Date Range: %d - %d", int(creationDateSliderMin.Value), int(creationDateSliderMax.Value)))
		firstAlbumLabel.SetText(fmt.Sprintf("First Album Date Range: %d - %d", int(firstAlbumSliderMin.Value), int(firstAlbumSliderMax.Value)))
		numberOfMembersLabel.SetText(fmt.Sprintf("Number of Members Range: %d - %d", int(numberOfMembersSliderMin.Value), int(numberOfMembersSliderMax.Value)))
	}

	creationDateSliderMin.OnChanged = func(value float64) {
		if creationDateSliderMin.Value >= creationDateSliderMax.Value {
			creationDateSliderMin.SetValue(creationDateSliderMax.Value)
		}
		updateLabels()
	}
	creationDateSliderMax.OnChanged = func(value float64) {
		if creationDateSliderMax.Value <= creationDateSliderMin.Value {
			creationDateSliderMax.SetValue(creationDateSliderMin.Value)
		}
		updateLabels()
	}

	firstAlbumSliderMin.OnChanged = func(value float64) {
		if firstAlbumSliderMin.Value >= firstAlbumSliderMax.Value {
			firstAlbumSliderMin.SetValue(firstAlbumSliderMax.Value)
		}
		updateLabels()
	}
	firstAlbumSliderMax.OnChanged = func(value float64) {
		if firstAlbumSliderMax.Value <= firstAlbumSliderMin.Value {
			firstAlbumSliderMax.SetValue(firstAlbumSliderMin.Value)
		}
		updateLabels()
	}

	numberOfMembersSliderMin.OnChanged = func(value float64) {
		if numberOfMembersSliderMin.Value >= numberOfMembersSliderMax.Value {
			numberOfMembersSliderMin.SetValue(numberOfMembersSliderMax.Value)
		}
		updateLabels()
	}
	numberOfMembersSliderMax.OnChanged = func(value float64) {
		if numberOfMembersSliderMax.Value <= numberOfMembersSliderMin.Value {
			numberOfMembersSliderMax.SetValue(numberOfMembersSliderMin.Value)
		}
		updateLabels()
	}

	creationDateContainer := container.NewVBox(creationDateLabel, creationDateSliderMin, creationDateSliderMax)
	firstAlbumContainer := container.NewVBox(widget.NewSeparator(), firstAlbumLabel, firstAlbumSliderMin, firstAlbumSliderMax)
	numberOfMembersContainer := container.NewVBox(widget.NewSeparator(), numberOfMembersLabel, numberOfMembersSliderMin, numberOfMembersSliderMax)

	locations := container.NewGridWithColumns(1)
	var locsCheck []*widget.Check
	for _, location := range core.LocationsCountry {
		locsCheck = append(locsCheck, widget.NewCheck(core.FirstLetterUpper(location), func(checked bool) {

		}))
	}
	for _, locsCheck := range locsCheck {
		locations.Add(locsCheck)
	}
	lab := widget.NewLabel(" \n \n \n \n \n \n \n")
	locationsContainer := container.NewVBox(widget.NewLabel("Locations"), container.NewHBox(lab, container.NewVScroll(locations)))

	applyButton := widget.NewButton("Apply Filters", func() {

	})

	resetButton := widget.NewButton("Reset Filters", func() {
		// Ici, vous pouvez maintenant accéder directement aux variables
		creationDateSliderMin.SetValue(float64(core.CreationDateRange[0]))
		creationDateSliderMax.SetValue(float64(core.CreationDateRange[1]))
		firstAlbumSliderMin.SetValue(float64(core.FirstAlbumDateRange[0]))
		firstAlbumSliderMax.SetValue(float64(core.FirstAlbumDateRange[1]))
		numberOfMembersSliderMin.SetValue(float64(core.NumberOfMembersRange[0]))
		numberOfMembersSliderMax.SetValue(float64(core.NumberOfMembersRange[1]))
		for _, box := range locsCheck {
			box.SetChecked(false)
		}
	})

	filterContainer := container.NewVBox(
		creationDateContainer,
		firstAlbumContainer,
		numberOfMembersContainer,
		locationsContainer,
		container.NewGridWithColumns(2, applyButton, resetButton),
	)

	topContainer := container.NewVBox(
		container.NewCenter(titleLabel),
		searchEntry,
		suggestionBox,
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
