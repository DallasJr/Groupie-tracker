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
	"strconv"
)

func LoadArtistPage(artist structs.Artist, myWindow fyne.Window) {
	homeButton := widget.NewButton("Home", func() {
		LoadMainPage(myWindow)
	})
	var favButton *widget.Button
	favButton = widget.NewButton("Favorite ", func() {
		if containsInt(core.Favorites, artist.ID) {
			core.RemoveFavorite(artist.ID)
		} else {
			core.AddFavorite(artist.ID)
		}
		setFavButtonIcon(favButton, artist.ID)
		err := core.SaveFavorites()
		if err != nil {
			return
		}
	})
	setFavButtonIcon(favButton, artist.ID)
	buttonContainer := container.NewGridWithColumns(2, container.NewVBox(homeButton), container.NewVBox(favButton))
	picture := artist.GetImage()
	picture.FillMode = canvas.ImageFillContain
	picture.SetMinSize(fyne.NewSize(150, 150))

	artistLabel := canvas.NewText(artist.Name, color.White)
	artistLabel.TextSize = 20

	title := container.NewGridWithColumns(4, layout.NewSpacer(), picture, container.NewCenter(artistLabel), layout.NewSpacer())

	membersList := widget.NewAccordion()
	groupContainer := container.NewVBox()
	for _, member := range artist.Members {
		memberButton := widget.NewButton(" - "+member, func() {})
		groupContainer.Add(memberButton)
	}
	membersItem := widget.NewAccordionItem("Members", groupContainer)
	membersItem.Open = true
	membersList.Append(membersItem)

	creation := canvas.NewText("Since "+strconv.Itoa(artist.CreationDate), color.White)
	creation.TextSize = 50
	firstAlbum := canvas.NewText("First Album: "+artist.GetFirstAlbum(), color.White)
	firstAlbum.TextSize = 50
	datesContainer := container.NewVBox(widget.NewSeparator(), container.NewCenter(creation), widget.NewSeparator(), container.NewCenter(firstAlbum), widget.NewSeparator())

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
	content := container.NewVScroll(container.NewVBox(buttonContainer, title, membersList, datesContainer, concertsContainer))
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
