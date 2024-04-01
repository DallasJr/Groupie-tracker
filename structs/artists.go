package structs

import (
	"fmt"
	"fyne.io/fyne/v2"
	"image"
	"net/http"
	"strings"

	"fyne.io/fyne/v2/canvas"
)

var Artists []Artist

type Artist struct {
	ID           int       `json:"id"`
	Image        string    `json:"image"`
	Name         string    `json:"name"`
	Members      []string  `json:"members"`
	CreationDate int       `json:"creationDate"`
	FirstAlbum   string    `json:"firstAlbum"`
	Locations    []string  `json:"-"`
	ConcertDates []string  `json:"-"`
	Relations    Relations `json:"-"`
}

var ImageArtist map[int]*canvas.Image

// On charge les images des artistes depuis l'API
// et les stock dans la map artist/image ImageArtist
func StoreArtistImage(artist Artist) {
	resp, err := http.Get(artist.Image)
	if err != nil {
		fmt.Println("Failed to load image:", err)
		return
	}
	defer resp.Body.Close()
	img, _, err := image.Decode(resp.Body)
	if err != nil {
		fmt.Println("Failed to decode image:", err)
		return
	}
	ImageArtist[artist.ID] = canvas.NewImageFromImage(img)
}

func (artist *Artist) GetImage() *canvas.Image {

	// Récupéré depuis la map artist/image
	formatted := canvas.NewImageFromImage(ImageArtist[artist.ID].Image)

	//Formattage
	formatted.FillMode = canvas.ImageFillContain
	formatted.SetMinSize(fyne.NewSize(100, 100))

	return formatted

}

func (artist *Artist) GetFirstAlbum() string {
	return GetFormattedDate(artist.FirstAlbum)
}

func GetFormattedDate(date string) string {
	return strings.ReplaceAll(date, "-", "/")
}

func GetArtist(id int) Artist {
	for _, artist := range Artists {
		if artist.ID == id {
			return artist
		}
	}
	return Artist{}
}
