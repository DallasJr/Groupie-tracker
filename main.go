package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Links struct {
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relation  string `json:"relation"`
}

func main() {
	// URL de l'API
	apiURL := "https://groupietrackers.herokuapp.com/api"

	// Faire une requête GET à l'API
	response, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Erreur lors de la requête :", err)
		return
	}
	defer response.Body.Close()

	// Structure pour stocker les liens
	var links Links

	// Décoder les données JSON
	err = json.NewDecoder(response.Body).Decode(&links)
	if err != nil {
		fmt.Println("Erreur lors du décodage des données JSON :", err)
		return
	}

	// Afficher les liens
	fmt.Println("Artists:", links.Artists)
	fmt.Println("Locations:", links.Locations)
	fmt.Println("Dates:", links.Dates)
	fmt.Println("Relation:", links.Relation)
}

func GetidArtists() {

}

func GetidLocation() {

}

func GetidRelation() {

}

func searchid() {

}
