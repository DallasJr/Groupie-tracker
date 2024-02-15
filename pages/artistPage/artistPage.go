package artistPage

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image"
	"net/http"
)

type OSMResponse struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

func LoadPage(myWindow fyne.Window) {
	url := "https://maptoolkit.p.rapidapi.com/staticmap?center=48.20835%2C16.3725&zoom=11&size=640x480&maptype=toursprung-terrain&format=png"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "a949d6c3c2mshcf7e90db304347cp17abc6jsn3307238f05d0")
	req.Header.Add("X-RapidAPI-Host", "maptoolkit.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	fmt.Println(res)
	img, _, err := image.Decode(res.Body)
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return
	}
	fyneImg := canvas.NewImageFromImage(img)

	// Create a container to hold the image
	content := container.NewMax(fyneImg)
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(800, 500))
}
