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
		if len(suggestions) > 10 {
			button := widget.NewButton("Too much results", func() {})
			suggestionBox.Add(button)
		} else {
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

	// Fonction pour effectuer une recherche
	performSearch := func() {

		// Supprimer les anciens résultats
		resultsContainer.RemoveAll()

		// Récuperer l'input
		searchInput := searchEntry.Text

		// Lancer la recherche
		searchResults = core.Search(searchInput)

		// Pour chaque résultat, on crée son objet
		for _, art := range searchResults {

			// On redéfini 'art' pour récuperer la bonne valeur de 'art'
			// pour éviter de se référer à la même variable art
			art := art

			// Récuperer et formatter l'image de l'artist
			picture := art.GetImage()

			artistLabel := widget.NewLabel(art.Name)

			// Le bouton rédirige vers le profile de l'artiste
			button := widget.NewButton("", func() {
				LoadArtistPage(art, myWindow)
			})

			// Le bouton est transparent de base et visible quand on a le curseur dessus
			button.Importance = widget.LowImportance

			button.SetIcon(theme.NavigateNextIcon())
			namepicture := container.NewGridWithColumns(2,
				picture,
				artistLabel,
			)

			// Le bouton est par dessus du container comportant l'image et le nom
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

	// Fonction qui s'execute quand l'utilisateur modifie dans la bar de recherche
	searchEntry.OnChanged = func(text string) {
		updateSuggestions(text)
		performSearch()
	}

	// Filters:
	// Déclaration et initialisation des sliders
	// Date de création
	creationDateSliderMin := widget.NewSlider(float64(core.CreationDateRange[0]), float64(core.CreationDateRange[1]))
	creationDateSliderMin.SetValue(core.CreationDateValue[0])
	creationDateSliderMax := widget.NewSlider(float64(core.CreationDateRange[0]), float64(core.CreationDateRange[1]))
	creationDateSliderMax.SetValue(core.CreationDateValue[1])

	// 1ère Album
	firstAlbumSliderMin := widget.NewSlider(float64(core.FirstAlbumDateRange[0]), float64(core.FirstAlbumDateRange[1]))
	firstAlbumSliderMin.SetValue(core.FirstAlbumDateValue[0])
	firstAlbumSliderMax := widget.NewSlider(float64(core.FirstAlbumDateRange[0]), float64(core.FirstAlbumDateRange[1]))
	firstAlbumSliderMax.SetValue(core.FirstAlbumDateValue[1])

	// Nombre de membres
	numberOfMembersSliderMin := widget.NewSlider(float64(core.NumberOfMembersRange[0]), float64(core.NumberOfMembersRange[1]))
	numberOfMembersSliderMin.SetValue(core.NumberOfMembersValue[0])
	numberOfMembersSliderMax := widget.NewSlider(float64(core.NumberOfMembersRange[0]), float64(core.NumberOfMembersRange[1]))
	numberOfMembersSliderMax.SetValue(core.NumberOfMembersValue[1])

	// Les labels avec les valeurs des sliders
	creationDateLabel := widget.NewLabel(fmt.Sprintf("Creation Date Range: %d - %d", int(creationDateSliderMin.Value), int(creationDateSliderMax.Value)))
	firstAlbumLabel := widget.NewLabel(fmt.Sprintf("First Album Date Range: %d - %d", int(firstAlbumSliderMin.Value), int(firstAlbumSliderMax.Value)))
	numberOfMembersLabel := widget.NewLabel(fmt.Sprintf("Number of Members Range: %d - %d", int(numberOfMembersSliderMin.Value), int(numberOfMembersSliderMax.Value)))

	locations := container.NewVBox()
	var locsCheck []*widget.Check

	// Fonction qui met à jour les valeurs pour le filtrage et des labels
	updateLabels := func() {
		creationDateLabel.SetText(fmt.Sprintf("Creation Date Range: %d - %d", int(creationDateSliderMin.Value), int(creationDateSliderMax.Value)))
		firstAlbumLabel.SetText(fmt.Sprintf("First Album Date Range: %d - %d", int(firstAlbumSliderMin.Value), int(firstAlbumSliderMax.Value)))
		numberOfMembersLabel.SetText(fmt.Sprintf("Number of Members Range: %d - %d", int(numberOfMembersSliderMin.Value), int(numberOfMembersSliderMax.Value)))
	}

	// Pour tout les pays possible, on crée un item checkbox
	for _, location := range core.LocationsCountry {
		name := core.FirstLetterUpper(location)
		check := widget.NewCheck(name, func(checked bool) {
			updateLabels()
		})

		// Si le pays était déjà cocher, on le recoche
		if core.ContainsString(core.LocationsCountryChecked, name) {
			check.SetChecked(true)
		}

		locsCheck = append(locsCheck, check)
	}

	updateLabels()

	// On ajoute tout les checkbox des pays dans le container 'locations'
	for _, locsCheck := range locsCheck {
		locations.Add(locsCheck)
	}
	lab := widget.NewLabel("\n\n\n\n\n\n\n")

	// Container final des locations pour les filtres
	locationsContainer := container.NewVBox(widget.NewLabel("Locations"), container.NewHBox(lab, container.NewVScroll(locations)))

	// Effectuer une recherche pour avoir des résultats au démarrage
	performSearch()

	// A chaque changement de valeurs dans sliders on lance la fonction 'updateLabels'
	creationDateSliderMin.OnChanged = func(value float64) {
		// Vérifie que la valeur min est <= à la valeur max, sinon on bloque
		if creationDateSliderMin.Value >= creationDateSliderMax.Value {
			creationDateSliderMin.SetValue(creationDateSliderMax.Value)
		}
		updateLabels()
	}
	creationDateSliderMax.OnChanged = func(value float64) {
		// Vérifie que la valeur min est <= à la valeur max, sinon on bloque
		if creationDateSliderMax.Value <= creationDateSliderMin.Value {
			creationDateSliderMax.SetValue(creationDateSliderMin.Value)
		}
		updateLabels()
	}

	firstAlbumSliderMin.OnChanged = func(value float64) {
		// Vérifie que la valeur min est <= à la valeur max, sinon on bloque
		if firstAlbumSliderMin.Value >= firstAlbumSliderMax.Value {
			firstAlbumSliderMin.SetValue(firstAlbumSliderMax.Value)
		}
		updateLabels()
	}
	firstAlbumSliderMax.OnChanged = func(value float64) {
		// Vérifie que la valeur min est <= à la valeur max, sinon on bloque
		if firstAlbumSliderMax.Value <= firstAlbumSliderMin.Value {
			firstAlbumSliderMax.SetValue(firstAlbumSliderMin.Value)
		}
		updateLabels()
	}

	numberOfMembersSliderMin.OnChanged = func(value float64) {
		// Vérifie que la valeur min est <= à la valeur max, sinon on bloque
		if numberOfMembersSliderMin.Value >= numberOfMembersSliderMax.Value {
			numberOfMembersSliderMin.SetValue(numberOfMembersSliderMax.Value)
		}
		updateLabels()
	}
	numberOfMembersSliderMax.OnChanged = func(value float64) {
		// Vérifie que la valeur min est <= à la valeur max, sinon on bloque
		if numberOfMembersSliderMax.Value <= numberOfMembersSliderMin.Value {
			numberOfMembersSliderMax.SetValue(numberOfMembersSliderMin.Value)
		}
		updateLabels()
	}

	creationDateContainer := container.NewVBox(creationDateLabel, creationDateSliderMin, creationDateSliderMax)
	firstAlbumContainer := container.NewVBox(widget.NewSeparator(), firstAlbumLabel, firstAlbumSliderMin, firstAlbumSliderMax)
	numberOfMembersContainer := container.NewVBox(widget.NewSeparator(), numberOfMembersLabel, numberOfMembersSliderMin, numberOfMembersSliderMax)

	// Apply enregistre les valeurs et lance une recherche
	applyButton := widget.NewButton("Apply Filters", func() {
		core.CreationDateValue[0] = creationDateSliderMin.Value
		core.CreationDateValue[1] = creationDateSliderMax.Value
		core.FirstAlbumDateValue[0] = firstAlbumSliderMin.Value
		core.FirstAlbumDateValue[1] = firstAlbumSliderMax.Value
		core.NumberOfMembersValue[0] = numberOfMembersSliderMin.Value
		core.NumberOfMembersValue[1] = numberOfMembersSliderMax.Value
		var checkedLocations []string
		for _, box := range locsCheck {
			if box.Checked {
				checkedLocations = append(checkedLocations, box.Text)
			}
		}
		core.LocationsCountryChecked = checkedLocations
		updateSuggestions(searchEntry.Text)
		performSearch()
	})

	// Bouton qui reset les valeurs des filtres et effectue une recherche automatiquement après
	resetButton := widget.NewButton("Reset Filters", func() {
		creationDateSliderMin.SetValue(float64(core.CreationDateRange[0]))
		creationDateSliderMax.SetValue(float64(core.CreationDateRange[1]))
		firstAlbumSliderMin.SetValue(float64(core.FirstAlbumDateRange[0]))
		firstAlbumSliderMax.SetValue(float64(core.FirstAlbumDateRange[1]))
		numberOfMembersSliderMin.SetValue(float64(core.NumberOfMembersRange[0]))
		numberOfMembersSliderMax.SetValue(float64(core.NumberOfMembersRange[1]))
		for _, box := range locsCheck {
			box.SetChecked(false)
		}
		core.CreationDateValue[0] = creationDateSliderMin.Value
		core.CreationDateValue[1] = creationDateSliderMax.Value
		core.FirstAlbumDateValue[0] = firstAlbumSliderMin.Value
		core.FirstAlbumDateValue[1] = firstAlbumSliderMax.Value
		core.NumberOfMembersValue[0] = numberOfMembersSliderMin.Value
		core.NumberOfMembersValue[1] = numberOfMembersSliderMax.Value
		core.LocationsCountryChecked = []string{}
		updateSuggestions(searchEntry.Text)
		performSearch()
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

	// Les artists favoris
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
