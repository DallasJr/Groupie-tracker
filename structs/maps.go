package structs

import (
	"encoding/json"
	"fmt"
	"image"
	"io"
	"net/http"
	"strconv"
	"strings"
)

var width = 300
var height = 300

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

func NewMap(center string, zoom int) *Map {
	parts := strings.Split(center, "-")
	city := strings.Replace(parts[0], "_", "%20", -1)
	country := strings.Replace(parts[1], "_", "%20", -1)
	url := "https://api.geoapify.com/v1/geocode/search?city=" + city + "&country=" + country + "&format=json&apiKey=d1ee7339aae647488e8b39534347dd95"
	res, err := http.Get(url)
	if err != nil {
		return nil
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	fmt.Println("Position Status: " + strconv.Itoa(res.StatusCode))
	var locRes LocResponse
	if err := json.Unmarshal(body, &locRes); err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	if len(locRes.Results) > 0 {
		lon := locRes.Results[0].Lon
		lat := locRes.Results[0].Lat
		fmt.Println("Longitude:", lon)
		fmt.Println("Latitude:", lat)
		return &Map{float32(lat), float32(lon), []Marker{}, zoom}
	} else {
		fmt.Println("No results found")
		return &Map{1.0, 1.0, []Marker{}, zoom}
	}
}

func (m *Map) GetImg() image.Image {
	url := m.GetURL()
	res, err := http.Get(url)
	if err != nil {
		return nil
	}
	defer res.Body.Close()
	fmt.Println("Map generation Status: " + strconv.Itoa(res.StatusCode))
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
	url += "&apiKey=d1ee7339aae647488e8b39534347dd95"
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
