package structs

import (
	"context"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
	"image"
	"net/http"
)

var ArtistsSpotifyDatas map[int]SpotifyData

var SpotifyClient spotify.Client

type SpotifyData struct {
	ID        spotify.ID
	TopTracks []Track
}

type Track struct {
	Name  string
	Image *canvas.Image
	Date  string
	URL   string
}

func InitializeSpotify() {

	logs := &clientcredentials.Config{
		ClientID:     "a3749c51d77549008226701463f3b294",
		ClientSecret: "5d03e0818f2c48609fa895ebbf38cc3c",
		TokenURL:     spotify.TokenURL,
	}
	client := logs.Client(context.Background())

	SpotifyClient = spotify.NewClient(client)
}

func LoadSpotifyData(artist Artist) {
	searchResults, err := SpotifyClient.Search(artist.Name, spotify.SearchTypeArtist)
	if err != nil {
		return
	}
	if len(searchResults.Artists.Artists) > 0 {
		artistID := searchResults.Artists.Artists[0].ID
		var tracks []Track
		topTracks, err := SpotifyClient.GetArtistsTopTracks(artistID, "US")
		if err != nil {
			return
		}
		for _, ttrack := range topTracks {
			tracks = append(tracks, Track{ttrack.Name, importImg(ttrack.Album.Images[0].URL), ttrack.Album.ReleaseDateTime().Format("02/01/2006"), ttrack.ExternalURLs["spotify"]})
		}
		ArtistsSpotifyDatas[artist.ID] = SpotifyData{artistID, tracks}

	} else {

	}
}

func importImg(url string) *canvas.Image {
	res, err := http.Get(url)
	if err != nil {
		return nil
	}
	defer res.Body.Close()
	img, _, err := image.Decode(res.Body)
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return nil
	}
	fyneImg := canvas.NewImageFromImage(img)
	fyneImg.FillMode = canvas.ImageFillContain
	fyneImg.SetMinSize(fyne.NewSize(250, 250))
	return fyneImg
}
