package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	// URL de l'API
	apiURL := "https://groupietrackers.herokuapp.com/api/artists"

	// Faire une requête GET à l'API
	response, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Erreur lors de la requête :", err)
		return
	}
	defer response.Body.Close()

	// Structure pour stocker les données JSON
	var data []map[string]interface{}

	// Décoder les données JSON directement depuis la réponse
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&data)
	if err != nil {
		fmt.Println("Erreur lors du décodage des données JSON :", err)
		return
	}

	// Afficher les données pour vérification
	fmt.Println(data)
}
