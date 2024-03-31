package core

import (
	"fmt"
	"groupie-tracker/structs"
	"strings"
)

// Recherche d'artistes générique
func Search(query string) []structs.Artist {
	var results []structs.Artist

	for _, artist := range structs.Artists {
		if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(query)) {
			results = append(results, artist)
			continue
		}
		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), strings.ToLower(query)) {
				results = append(results, artist)
				continue
			}
		}
		for _, loc := range artist.Locations {
			if strings.Contains(strings.ToLower(loc), strings.ToLower(query)) {
				results = append(results, artist)
				continue
			}
		}
		if strings.Contains(strings.ToLower(artist.FirstAlbum), strings.ToLower(strings.Replace(query, "/", "-", -1))) {
			results = append(results, artist)
			continue
		}
		if strings.Contains(strings.ToLower(fmt.Sprint(artist.CreationDate)), strings.ToLower(query)) {
			results = append(results, artist)
			continue
		}
		for _, location := range artist.Locations {
			ok := false
			for _, words := range strings.Fields(query) {
				if strings.Contains(strings.ToLower(location), strings.ToLower(words)) {
					ok = true
				} else {
					ok = false
					break
				}
			}
			if ok {
				results = append(results, artist)
				break
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
	if query == "" {
		return suggestions
	}

	for _, artist := range structs.Artists {
		matches := make(map[string]bool) // Pour garder une trace des types de correspondances trouvées

		if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(query)) {
			matches["Artist"] = true
		}

		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), strings.ToLower(query)) {
				matches["Member"] = true
				break
			}
		}

		for _, location := range artist.Locations {
			ok := false
			for _, words := range strings.Fields(query) {
				if strings.Contains(strings.ToLower(location), strings.ToLower(words)) {
					ok = true
				} else {
					ok = false
					break
				}
			}
			if ok {
				matches["Concert Dates"] = true
				break
			}
		}

		if strings.Contains(strings.ToLower(artist.FirstAlbum), strings.ToLower(strings.Replace(query, "/", "-", -1))) {
			matches["First Album"] = true
		}

		if strings.Contains(strings.ToLower(fmt.Sprint(artist.CreationDate)), strings.ToLower(query)) {
			matches["Creation Date"] = true
		}

		if len(matches) > 0 {
			var matchTypes []string
			for match := range matches {
				matchTypes = append(matchTypes, match)
			}
			suggestions = append(suggestions, fmt.Sprintf("%s - %s", artist.Name, strings.Join(matchTypes, "/")))
		}
	}

	return suggestions
}
