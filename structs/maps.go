package structs

import (
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"image"
	"io"
	"net/http"
	"strconv"
	"strings"
)

const width = 300
const height = 300
const zoom = 4
const apiKey = "b4110a75f3ed466d8ee295c29da92f87"
const Disabled = true

var ImageMap map[string]*canvas.Image

type Map struct {
	lat     float32
	long    float32
	markers []Marker
	zoom    int
}

type LocResponse struct {
	Results []struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"results"`
}

func GenerateMapImage(location string) {
	m := NewMap(location)
	m.AddMarker(m.GetLat(), m.GetLong(), "394e70", "users")
	img := m.GetImg()
	fyneImg := canvas.NewImageFromImage(img)
	fyneImg.FillMode = canvas.ImageFillContain
	fyneImg.SetMinSize(fyne.NewSize(300, 300))
	ImageMap[location] = fyneImg
}

func GetMapImage(location string) *canvas.Image {
	if img, ok := ImageMap[location]; ok {
		return img
	}
	blankImg := canvas.NewImageFromResource(nil)
	return blankImg
}

func NewMap(center string) *Map {
	parts := strings.Split(center, "-")
	city := strings.Replace(parts[0], "_", "%20", -1)
	country := strings.Replace(parts[1], "_", "%20", -1)
	url := "https://api.geoapify.com/v1/geocode/search?city=" + city + "&country=" + country + "&format=json&apiKey=" + apiKey
	res, err := http.Get(url)
	if err != nil {
		return nil
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	//fmt.Println("Position Status: " + strconv.Itoa(res.StatusCode))
	var locRes LocResponse
	if err := json.Unmarshal(body, &locRes); err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	if len(locRes.Results) > 0 {
		lon := locRes.Results[0].Lon
		lat := locRes.Results[0].Lat
		//fmt.Println("Longitude:", lon)
		//fmt.Println("Latitude:", lat)
		return &Map{float32(lat), float32(lon), []Marker{}, zoom}
	} else {
		fmt.Println("No results found")
		return &Map{1.0, 1.0, []Marker{}, zoom}
	}
}

func GetFormattedLocationName(location string) (string, string) {
	parts := strings.Split(location, "-")
	city := strings.Replace(parts[0], "_", " ", -1)
	country := strings.Replace(parts[1], "_", " ", -1)
	words := strings.Fields(city)
	for i, word := range words {
		words[i] = strings.ToUpper(string(word[0])) + word[1:]
	}
	city = strings.Join(words, " ")
	words = strings.Fields(country)
	for i, word := range words {
		words[i] = strings.ToUpper(string(word[0])) + word[1:]
	}
	country = strings.Join(words, " ")
	return city, country
}

func (m *Map) GetImg() image.Image {
	url := m.GetURL()
	res, err := http.Get(url)
	if err != nil {
		return nil
	}
	defer res.Body.Close()
	//fmt.Println("Map generation Status: " + strconv.Itoa(res.StatusCode))
	img, _, err := image.Decode(res.Body)
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return nil
	}
	return img
}

func (m *Map) GetURL() string {
	url := "https://maps.geoapify.com/v1/staticmap?style=maptiler-3d"
	url += "&width=" + strconv.Itoa(width) + "&height=" + strconv.Itoa(height)
	url += "&center=lonlat:" + toString(m.long) + "," + toString(m.lat)
	url += "&zoom=" + strconv.Itoa(m.zoom)
	if len(m.markers) > 0 {
		url += "&marker="
		for i, marker := range m.markers {
			url += "lonlat:" + toString(marker.long) + "," + toString(marker.lat)
			url += ";type:awesome"
			url += ";color:%23" + marker.colorHex
			url += ";size:x-large"
			url += ";icon:" + marker.icon
			if i != len(m.markers)-1 {
				url += "|"
			}
		}
	}
	url += "&scaleFactor=2"
	url += "&apiKey=" + apiKey
	return url
}

type Marker struct {
	lat      float32
	long     float32
	colorHex string
	icon     string
	//icons here: https://fontawesome.com/v5/search?o=r&m=free&s=solid
}

func (m *Map) AddMarker(lat float32, long float32, color string, icon string) {
	m.markers = append(m.markers, Marker{lat, long, color, icon})
}

func (m *Map) GetLat() (lat float32) {
	return m.lat
}

func (m *Map) GetLong() (long float32) {
	return m.long
}

func toString(val float32) string {
	return strconv.FormatFloat(float64(val), 'f', -1, 32)
}
