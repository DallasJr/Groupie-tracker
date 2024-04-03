package pages

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"groupie-tracker/core"
	"groupie-tracker/structs"
	"image/color"
	"net/url"
	"strconv"
)

func LoadArtistPage(myApp fyne.App, artist structs.Artist, myWindow fyne.Window) {
	homeButton := widget.NewButton("Home", func() {
		LoadMainPage(myApp, myWindow)
	})

	var favButton *widget.Button
	favButton = widget.NewButton("Favorite ", func() {

		// Ajoute ou retire l'artiste des favoris
		if containsInt(core.Favorites, artist.ID) {
			core.RemoveFavorite(artist.ID)
		} else {
			core.AddFavorite(artist.ID)
		}

		// Refresh l'icon du bouton favoris
		setFavButtonIcon(favButton, artist.ID)

		// Sauvegarde dans le fichier
		err := core.SaveFavorites()
		if err != nil {
			return
		}

	})
	// Refresh l'icon du bouton favoris
	setFavButtonIcon(favButton, artist.ID)

	buttonContainer := container.NewGridWithColumns(2, container.NewVBox(homeButton), container.NewVBox(favButton))

	picture := artist.GetImage()
	picture.FillMode = canvas.ImageFillContain
	picture.SetMinSize(fyne.NewSize(300, 300))

	artistLabel := canvas.NewText(artist.Name, color.White)
	artistLabel.TextStyle = fyne.TextStyle{Bold: true}
	artistLabel.TextSize = 30

	title := container.NewGridWithColumns(4, layout.NewSpacer(), picture, container.NewCenter(artistLabel), layout.NewSpacer())

	// Membres
	membersList := widget.NewAccordion()
	groupContainer := container.NewVBox()
	for _, member := range artist.Members {
		memberButton := widget.NewButton(" - "+member, func() {})
		groupContainer.Add(memberButton)
	}
	membersItem := widget.NewAccordionItem("Members", groupContainer)
	membersItem.Open = true
	membersList.Append(membersItem)

	// Date de création
	creation := canvas.NewText("Since "+strconv.Itoa(artist.CreationDate), color.White)
	creation.TextSize = 50

	// Date du 1er album
	firstAlbum := canvas.NewText("First Album: "+artist.GetFirstAlbum(), color.White)
	firstAlbum.TextSize = 50

	datesContainer := container.NewVBox(widget.NewSeparator(), container.NewCenter(creation), widget.NewSeparator(), container.NewCenter(firstAlbum), widget.NewSeparator())

	// Affichage des concerts (lieu + dates)
	concertsContainer := container.NewVBox()
	concertsLabel := canvas.NewText("Concerts:", color.White)
	concertsLabel.TextSize = 50
	concertsContainer.Add(container.NewCenter(concertsLabel))
	if len(artist.Locations) > 0 {
		concerts := container.NewGridWithColumns(3)
		for location, dates := range artist.Relations.DatesLocations {
			concertCard := container.NewVBox()
			concertCard.Add(structs.GetMapImage(location))
			city, country := structs.GetFormattedLocationName(location)
			locationText := canvas.NewText(city+" "+country, color.White)
			locationText.TextStyle = fyne.TextStyle{Bold: true}
			locationText.TextSize = 20
			concertCard.Add(container.NewCenter(locationText))
			datesLabel := widget.NewLabel("Dates:")
			concertCard.Add(container.NewCenter(datesLabel))

			for _, date := range dates {
				concertCard.Add(container.NewCenter(widget.NewLabel(structs.GetFormattedDate(date))))
			}
			concerts.Add(concertCard)
		}
		concertsContainer.Add(concerts)
		concertsContainer.Add(widget.NewSeparator())
	} else {
		noConcerts := canvas.NewText("No concerts", color.White)
		noConcerts.TextSize = 20
		concertsContainer.Add(noConcerts)
		concertsContainer.Add(widget.NewSeparator())
	}
	topTracksContainer := container.NewVBox()
	topTracksLabel := canvas.NewText("Top Tracks:", color.White)
	topTracksLabel.TextSize = 50
	topTracksContainer.Add(container.NewCenter(topTracksLabel))
	spotifyLabel := canvas.NewText("from Spotify", color.CMYK{C: 80, Y: 80})
	spotifyLabel.TextSize = 20
	topTracksContainer.Add(container.NewCenter(spotifyLabel))
	if len(structs.ArtistsSpotifyDatas[artist.ID].TopTracks) > 0 {
		tracks := container.NewGridWithColumns(2)
		for _, track := range structs.ArtistsSpotifyDatas[artist.ID].TopTracks {
			nameText := canvas.NewText(track.Name, color.White)
			nameText.TextStyle = fyne.TextStyle{Bold: true}
			nameText.TextSize = 20
			indexText := canvas.NewText(strconv.Itoa(len(tracks.Objects)+1), color.White)
			indexText.TextStyle = fyne.TextStyle{Bold: true}
			indexText.TextSize = 18
			dateText := canvas.NewText(track.Date, color.White)
			dateText.TextSize = 18
			card := widget.NewCard("", "", container.NewVBox(container.NewCenter(indexText), track.Image, container.NewCenter(nameText), container.NewCenter(dateText)))
			button := widget.NewButton("", func() {
				u, _ := url.Parse(track.URL)
				_ = myApp.OpenURL(u)
			})
			button.Importance = widget.LowImportance
			button.SetIcon(theme.NavigateNextIcon())
			finalCard := container.New(layout.NewBorderLayout(nil, nil, nil, nil),
				card,
				container.NewGridWithColumns(3, layout.NewSpacer(), button, layout.NewSpacer()))
			tracks.Add(finalCard)
		}
		topTracksContainer.Add(tracks)
		topTracksContainer.Add(widget.NewSeparator())
	} else {
		noTracks := canvas.NewText("No tracks", color.White)
		noTracks.TextSize = 20
		topTracksContainer.Add(noTracks)
		topTracksContainer.Add(widget.NewSeparator())
	}

	content := container.NewVScroll(container.NewVBox(buttonContainer, title, membersList, datesContainer, concertsContainer, topTracksContainer))
	myWindow.SetContent(content)
}

func setFavButtonIcon(btn *widget.Button, artistID int) {
	if containsInt(core.Favorites, artistID) {
		btn.SetIcon(theme.CheckButtonCheckedIcon())
	} else {
		btn.SetIcon(theme.CheckButtonIcon())
	}
}

func containsInt(arr []int, str int) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}
