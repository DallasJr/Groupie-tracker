package core

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/structs"
	"net/http"
)

// LoadAll charge toutes les données à partir des liens donnés.
func Load() error {
	// Charger les artistes
	links := map[string]string{
		"artists":   "https://groupietrackers.herokuapp.com/api/artists",
		"locations": "https://groupietrackers.herokuapp.com/api/locations",
		"dates":     "https://groupietrackers.herokuapp.com/api/dates",
		"relations": "https://groupietrackers.herokuapp.com/api/relation", // Correction de la clé ici
	}

	if err := LoadArtists(links["artists"]); err != nil {
		return fmt.Errorf("erreur lors du chargement des artistes : %v", err)
	}

	// Charger les localisations
	if err := LoadLocations(links["locations"]); err != nil {
		return fmt.Errorf("erreur lors du chargement des localisations : %v", err)
	}

	// Charger les dates
	if err := LoadDates(links["dates"]); err != nil {
		return fmt.Errorf("erreur lors du chargement des dates : %v", err)
	}

	// Charger les relations
	if err := LoadRelations(links["relations"]); err != nil {
		return fmt.Errorf("erreur lors du chargement des relations : %v", err)
	}

	fmt.Println("Loaded artists: ")
	for _, artist := range structs.Artists {
		fmt.Print("- ")
		fmt.Println(artist.CreationDate)
	}

	return nil
}

// LoadArtists charge les artistes depuis l'API.
func LoadArtists(url string) error {
	fmt.Println("Loading artists ...")
	// Faire une requête GET à l'URL de l'API
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Vérifier le code de statut de la réponse
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("la requête a retourné un code de statut non-200 : %d", response.StatusCode)
	}

	// Décoder les données JSON
	err = json.NewDecoder(response.Body).Decode(&structs.Artists)
	if err != nil {
		return fmt.Errorf("erreur lors du décodage des données JSON : %v", err)
	}

	return nil
}

// LoadLocations charge les localisations depuis l'API.
func LoadLocations(url string) error {
	fmt.Println("Loading locations...")
	// Faire une requête GET à l'URL de l'API
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Vérifier le code de statut de la réponse
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("la requête a retourné un code de statut non-200 : %d", response.StatusCode)
	}

	// Décode les données JSON
	var locations structs.Locations
	if err := json.NewDecoder(response.Body).Decode(&locations); err != nil {
		return fmt.Errorf("erreur lors du décodage des données JSON : %v", err)
	}

	// Mettez à jour les données dans la structure globale ou faites ce qui est nécessaire
	// structs.Locations = locations

	return nil
}

// LoadDates charge les dates depuis l'API.
func LoadDates(url string) error {
	fmt.Println("Loading dates...")
	// Faire une requête GET à l'URL de l'API
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Vérifier le code de statut de la réponse
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("la requête a retourné un code de statut non-200 : %d", response.StatusCode)
	}

	// Décode les données JSON
	var dates structs.Dates
	if err := json.NewDecoder(response.Body).Decode(&dates); err != nil {
		return fmt.Errorf("erreur lors du décodage des données JSON : %v", err)
	}

	// Mettez à jour les données dans la structure globale ou faites ce qui est nécessaire
	// structs.Dates = dates

	return nil
}

// LoadRelations charge les relations depuis l'API.
func LoadRelations(url string) error {
	fmt.Println("Loading relations...")
	// Faire une requête GET à l'URL de l'API
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Vérifier le code de statut de la réponse
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("la requête a retourné un code de statut non-200 : %d", response.StatusCode)
	}

	// Décode les données JSON
	var relations structs.Relation
	if err := json.NewDecoder(response.Body).Decode(&relations); err != nil {
		return fmt.Errorf("erreur lors du décodage des données JSON : %v", err)
	}

	// Mettez à jour les données dans la structure globale ou faites ce qui est nécessaire
	// structs.Relation = relation

	return nil
}
