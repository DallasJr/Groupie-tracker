package structs

import (
	"fmt"
	"image"
	"net/http"

	"fyne.io/fyne/v2/canvas"
)

var Artists []Artist

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    []string `json:"-"`
	ConcertDates []string `json:"-"`
	Relations    string   `json:"-"`
}

func (artist *Artist) GetImage() *canvas.Image {
	resp, err := http.Get(artist.Image)
	if err != nil {
		fmt.Println("Failed to load image:", err)
		return nil
	}
	defer resp.Body.Close()

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		fmt.Println("Failed to decode image:", err)
		return nil
	}

	return canvas.NewImageFromImage(img)
}
