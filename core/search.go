package core

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/structs"
	"net/http"
	"strconv"
	"strings"
)

// refaire le load pour qu'il soit dinamique avec ce lien
func Load() {

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
		fmt.Println("- " + artist.Name)
	}
}

// Recherche d'artistes par nom
func SearchArtistsByName(query string) []structs.Artist {
	Load()
	var results []structs.Artist
	for _, artist := range structs.Artists {
		if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(query)) {
			results = append(results, artist)
		}
	}
	return results
}

// Recherche d'artistes par membre
func SearchArtistsByMember(query string) []structs.Artist {
	var results []structs.Artist
	for _, artist := range structs.Artists {
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
func SearchArtistsByCreationYear(year int) []structs.Artist {
	var results []structs.Artist
	for _, artist := range structs.Artists {
		if artist.CreationDate == year {
			results = append(results, artist)
		}
	}
	return results
}

// Recherche d'artistes générique
func Search(query string) []structs.Artist {
	var results []structs.Artist

	// Recherche par nom d'artiste
	nameResults := SearchArtistsByName(query)
	results = append(results, nameResults...)

	// Recherche par membre
	memberResults := SearchArtistsByMember(query)
	results = append(results, memberResults...)

	// Recherche par année de création (si la requête est un nombre)
	if year, err := strconv.Atoi(query); err == nil {
		yearResults := SearchArtistsByCreationYear(year)
		results = append(results, yearResults...)
	}

	// Recherche par d'autres champs comme les localisations, les dates de concert, etc.
	for _, artist := range structs.Artists {
		if strings.Contains(strings.ToLower(artist.Locations), strings.ToLower(query)) ||
			strings.Contains(strings.ToLower(artist.ConcertDates), strings.ToLower(query)) ||
			strings.Contains(strings.ToLower(artist.Relations), strings.ToLower(query)) {
			results = append(results, artist)
		}
	}

	// Supprimer les doublons potentiels
	results = removeDuplicates(results)

	return results
}

// Fonction utilitaire pour supprimer les doublons dans la liste des artistes
func removeDuplicates(artists []structs.Artist) []structs.Artist {
	encountered := map[int]bool{}
	var result []structs.Artist

	for _, artist := range artists {
		if !encountered[artist.ID] {
			encountered[artist.ID] = true
			result = append(result, artist)
		}
	}

	return result
}

func GetSuggestions(query string) []string {
	var suggestions []string

	for _, artist := range structs.Artists {
		// Vérifier si le nom de l'artiste contient la chaîne de caractères de la requête
		if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(query)) {
			suggestions = append(suggestions, artist.Name)
		}
	}

	return suggestions
}

//func GetContain(query string) string {
//	container.New()
//}
