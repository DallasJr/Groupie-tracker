package core

import (
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2/canvas"
	"groupie-tracker/structs"
	"io"
	"net/http"
	"os"
	"sort"
	"strings"
)

func Load() {
	// Load favorite artists from local data file
	LoadFavorites()

	// Load the artists from API
	LoadArtists()

	// Load artists concerts locations and dates from API
	loadLocations()
	loadDate()
	loadRelations()
}

const FavFile = "favs.json"

func LoadFavorites() {
	fmt.Println("Loading Favorites")

	// Vérifier l'existance du fichier .json
	if _, err := os.Stat(FavFile); err == nil {

		// Ouvrir et lire le fichier
		file, err := os.Open(FavFile)
		if err != nil {
			return
		}
		defer file.Close()
		jsonBytes, err := io.ReadAll(file)
		if err != nil {
			return
		}

		// Déccoder les données JSON et stocker dans Favorites
		err = json.Unmarshal(jsonBytes, &Favorites)
		if err != nil {
			return
		}
		fmt.Println("Loaded favorites file")
		return

	} else if os.IsNotExist(err) {
		fmt.Println("Data file doesn't exist yet")
	} else {
		fmt.Println("Error occurred while checking file existence:", err)
	}
}

func LoadArtists() {
	fmt.Println("Loading Artists")
	structs.ImageArtist = make(map[int]*canvas.Image)
	URL := "https://groupietrackers.herokuapp.com/api/artists"

	// Faire une requête GET à l'API
	response, err := http.Get(URL)
	if err != nil {
		fmt.Println("Erreur lors de la requête :", err)
		return
	}
	defer response.Body.Close()

	// Vérifier le code de statut de la réponse
	if response.StatusCode != http.StatusOK {
		fmt.Println("La requête a retourné un code de statut non-200 :", response.StatusCode)
		return
	}

	// Décoder les données JSON
	err = json.NewDecoder(response.Body).Decode(&structs.Artists)
	if err != nil {
		fmt.Println("Erreur lors du décodage des données JSON :", err)
		return
	}

	fmt.Println("Loading artists images and spotify data")
	structs.ArtistsSpotifyDatas = make(map[int]structs.SpotifyData)
	for _, artist := range structs.Artists {
		// Récuperer les valeurs min et max
		InitializeFiltersValues(artist, true)

		go structs.LoadSpotifyData(artist)
		// On charge les images des artistes et les stock dans la map artist/image ImageArtist
		go structs.StoreArtistImage(artist)

	}

	// Défini les valeurs par défaut
	InitializeFiltersValues(structs.Artist{}, false)
}

func loadLocations() {
	fmt.Println("Loading Locations")
	structs.ImageMap = make(map[string]*canvas.Image)
	var allLocations []string
	for i := range structs.Artists {
		url := "https://groupietrackers.herokuapp.com/api/locations/" + fmt.Sprint(structs.Artists[i].ID)
		response, err := http.Get(url)
		if err != nil {
			fmt.Println("Erreur lors de la requête :", err)
			return
		}
		defer response.Body.Close()

		// Vérifier le code de statut de la réponse
		if response.StatusCode != http.StatusOK {
			fmt.Println("La requête a retourné un code de statut non-200 :", response.StatusCode)
			return
		}
		// Décoder les données JSON
		var locations structs.Locations
		err = json.NewDecoder(response.Body).Decode(&locations)
		if err != nil {
			fmt.Println("Erreur lors du décodage des données JSON :", err)
			return
		}
		for _, loc := range locations.Locations {
			if !ContainsString(allLocations, loc) {
				allLocations = append(allLocations, loc)
			}
		}

		// Assigner les locations à l'artiste correspondant
		structs.Artists[i].Locations = locations.Locations
	}

	fmt.Println("Loading Map images")
	for _, loc := range allLocations {

		// Récupère et format les noms des pays pour les checkbox des filtres
		_, country := structs.GetFormattedLocationName(loc)
		if !ContainsString(LocationsCountry, strings.ToLower(country)) {
			LocationsCountry = append(LocationsCountry, strings.ToLower(country))
		}

		// Charge les images de map et les stock dans la map artist/image ImageMap
		go structs.GenerateMapImage(loc)
	}

	// Trie par ordre alphabétique
	sort.Strings(LocationsCountry)
}

func loadDate() {
	fmt.Println("Loading Dates")
	for i := range structs.Artists {
		url := "https://groupietrackers.herokuapp.com/api/dates/" + fmt.Sprint(structs.Artists[i].ID)
		response, err := http.Get(url)
		if err != nil {
			fmt.Println("Erreur lors de la requête :", err)
			return
		}
		defer response.Body.Close()

		// Vérifier le code de statut de la réponse
		if response.StatusCode != http.StatusOK {
			fmt.Println("La requête a retourné un code de statut non-200 :", response.StatusCode)
			return
		}
		// Décoder les données JSON
		var date structs.Dates
		err = json.NewDecoder(response.Body).Decode(&date)
		if err != nil {
			fmt.Println("Erreur lors du décodage des données JSON :", err)
			return
		}

		structs.Artists[i].ConcertDates = date.Date
	}
}

func loadRelations() {
	fmt.Println("Loading Relations")
	for i := range structs.Artists {
		url := "https://groupietrackers.herokuapp.com/api/relation/" + fmt.Sprint(structs.Artists[i].ID)
		response, err := http.Get(url)
		if err != nil {
			fmt.Println("Erreur lors de la requête :", err)
			return
		}
		defer response.Body.Close()

		// Vérifier le code de statut de la réponse
		if response.StatusCode != http.StatusOK {
			fmt.Println("La requête a retourné un code de statut non-200 :", response.StatusCode)
			return
		}
		// Décoder les données JSON
		var concerts structs.Relations
		err = json.NewDecoder(response.Body).Decode(&concerts)
		if err != nil {
			fmt.Println("Erreur lors du décodage des données JSON :", err)
			return
		}

		structs.Artists[i].Relations = concerts
	}
}
