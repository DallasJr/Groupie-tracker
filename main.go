package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

var artists []Artist

func main() {
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
	err = json.NewDecoder(response.Body).Decode(&artists)
	if err != nil {
		fmt.Println("Erreur lors du décodage des données JSON :", err)
		return
	}

	// Exemples de recherches
	fmt.Println("Recherche par nom:")
	fmt.Println(SearchArtistsByName("Queen"))

	fmt.Println("\nRecherche par membre:")
	fmt.Println(SearchArtistsByMember("Freddie Mercury"))

	fmt.Println("\nRecherche par année de création:")
	fmt.Println(SearchArtistsByCreationYear(1970))
}

// Recherche d'artistes par nom
func SearchArtistsByName(query string) []Artist {
	var results []Artist
	for _, artist := range artists {
		if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(query)) {
			results = append(results, artist)
		}
	}
	return results
}

// Recherche d'artistes par membre
func SearchArtistsByMember(query string) []Artist {
	var results []Artist
	for _, artist := range artists {
		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), strings.ToLower(query)) {
				results = append(results, artist)
				break
			}
		}
	}
	return results
}

// Recherche d'artistes par année de création
func SearchArtistsByCreationYear(year int) []Artist {
	var results []Artist
	for _, artist := range artists {
		if artist.CreationDate == year {
			results = append(results, artist)
		}
	}
	return results
}
