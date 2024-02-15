package search

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

var artist []Artist

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
	err = json.NewDecoder(response.Body).Decode(&artist)
	if err != nil {
		fmt.Println("Erreur lors du décodage des données JSON :", err)
		return
	}
	fmt.Println("Recherche par nom:")
	searchByName("Queen")

	fmt.Println("\nRecherche par membre:")
	searchByMember("Freddie Mercury")

	fmt.Println("\nRecherche par année de création:")
	searchByCreationYear(1970)
}

func searchByName(query string) {
	results := SearchArtistsByName(query)
	for _, artist := range results {
		printArtistInfo(artist)
	}
}

func searchByMember(query string) {
	results := SearchArtistsByMember(query)
	for _, artist := range results {
		printArtistInfo(artist)
	}
}

func searchByCreationYear(year int) {
	results := SearchArtistsByCreationYear(year)
	for _, artist := range results {
		printArtistInfo(artist)
	}
}

func printArtistInfo(artist Artist) {
	fmt.Println("ID:", artist.ID)
	fmt.Println("Image:", artist.Image)
	fmt.Println("Nom:", artist.Name)
	fmt.Println("Membres:", strings.Join(artist.Members, ", "))
	fmt.Println("Date de création:", artist.CreationDate)
	fmt.Println("Premier album:", artist.FirstAlbum)
}

func getArtistInfo(url string) {
	data, err := fetchData(url)
	if err != nil {
		fmt.Println("Erreur lors de la récupération des données depuis", url, ":", err)
		return
	}
	fmt.Println(data)
}

func fetchData(url string) (interface{}, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var data interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Recherche d'artistes par nom
func SearchArtistsByName(query string) []Artist {
	var results []Artist
	for _, artist := range artist {
		if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(query)) {
			results = append(results, artist)
		}
	}
	return results
}

// Recherche d'artistes par membre
func SearchArtistsByMember(query string) []Artist {
	var results []Artist
	for _, artist := range artist {
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
	for _, artist := range artist {
		if artist.CreationDate == year {
			results = append(results, artist)
		}
	}
	return results
}
