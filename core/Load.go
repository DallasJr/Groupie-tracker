package core

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/structs"
	"net/http"
)

func Load() {
	LoadArtists()
	loadLocations()
}
func LoadArtists() {
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

	fmt.Println("Loaded artists: ")
	for _, artist := range structs.Artists {
		for _, truc := range artist.Locations {
			fmt.Println("- " + truc)
		}
	}
}

func loadLocations() {
	for _, artist := range structs.Artists {
		url := "https://groupietrackers.herokuapp.com/api/locations/" + fmt.Sprint(artist.ID)
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
		artist.Locations = locations.Locations
	}
}
