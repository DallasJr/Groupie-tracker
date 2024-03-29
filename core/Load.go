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
	loadDate()
	loadRelations()
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
}

func loadLocations() {
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

		// Assigner les locations à l'artiste correspondant
		structs.Artists[i].Locations = locations.Locations
	}
}

func loadDate() {
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
		var relation structs.Relation
		err = json.NewDecoder(response.Body).Decode(&relation)
		if err != nil {
			fmt.Println("Erreur lors du décodage des données JSON :", err)
			return
		}

		structs.Artists[i].Relations.Relationlocation = relation.Relationlocation
		structs.Artists[i].Relations.Table_Dates = relation.Table_Dates
	}
}
