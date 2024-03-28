package core

import (
	"groupie-tracker/structs"
	"strconv"
	"strings"
)

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

	for _, artist := range structs.Artists {
		for loc := range artist.Locations {
			if strings.Contains(strings.ToLower(artist.Locations[loc]), strings.ToLower(query)) {
				results = append(results, artist)
			}
		}
		for concertdate := range artist.ConcertDates {
			if strings.Contains(strings.ToLower(artist.ConcertDates[concertdate]), strings.ToLower(query)) {
				results = append(results, artist)
			}
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
