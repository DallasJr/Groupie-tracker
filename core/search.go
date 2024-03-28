package core

import (
	"groupie-tracker/structs"
	"strings"
)

// Recherche d'artistes par nom
func SearchArtistsByName(query string) []structs.Artist {
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

// Recherche d'artistes par localisation
func SearchArtistsByLocation(location string) []structs.Artist {
	var results []structs.Artist

	for _, artist := range structs.Artists {
		if strings.Contains(strings.ToLower(strings.Join(artist.Locations.Locations, " ")), strings.ToLower(location)) {
			results = append(results, artist)
		}
	}

	return results
}

// Recherche d'artistes par date
func SearchArtistsByDate(date string) []structs.Artist {
	var results []structs.Artist

	for _, artist := range structs.Artists {
		if strings.Contains(strings.ToLower(strings.Join(artist.Locations.Dates.Dates, " ")), strings.ToLower(date)) {
			results = append(results, artist)
		}
	}

	return results
}

// Recherche d'artistes par relation
func SearchArtistsByRelation(relation string) []structs.Artist {
	var results []structs.Artist

	for _, artist := range structs.Artists {
		for _, datesLocation := range artist.Relations.Dates_Locations {
			if strings.Contains(strings.ToLower(strings.Join(datesLocation.Table_Dates, " ")), strings.ToLower(relation)) {
				results = append(results, artist)
				break
			}
		}
	}

	return results
}

// Recherche d'artistes générique
func Search(query string) []structs.Artist {
	Load()
	var results []structs.Artist

	// Recherche par nom d'artiste
	nameResults := SearchArtistsByName(query)
	results = append(results, nameResults...)

	// Recherche par membre
	memberResults := SearchArtistsByMember(query)
	results = append(results, memberResults...)

	// Recherche par localisation
	locationResults := SearchArtistsByLocation(query)
	results = append(results, locationResults...)

	// Recherche par date
	dateResults := SearchArtistsByDate(query)
	results = append(results, dateResults...)

	// Recherche par relation
	relationResults := SearchArtistsByRelation(query)
	results = append(results, relationResults...)

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
