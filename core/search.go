package core

import (
	"fmt"
	"groupie-tracker/structs"
	"regexp"
	"strconv"
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

// Recherche d'artistes par année du premiere album
func SearchArtistsByFirstAlbumYear(year int) []structs.Artist {
	var results []structs.Artist
	for _, artist := range structs.Artists {
		if getYearFromDate(artist.FirstAlbum) == year {
			results = append(results, artist)
		}
	}
	return results
}

// Recherche d'artistes par date du premiere album
func SearchArtistsByFirstAlbumDate(date string) []structs.Artist {
	var results []structs.Artist
	for _, artist := range structs.Artists {
		if artist.FirstAlbum == strings.Replace(date, "/", "-", -1) {
			results = append(results, artist)
		}
	}
	return results
}

// Recherche d'artistes générique
func Search(query string) []structs.Artist {
	//Load()
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

	// Recherche par année de première album
	if year, err := strconv.Atoi(query); err == nil {
		yearResults := SearchArtistsByFirstAlbumYear(year)
		results = append(results, yearResults...)
	}

	// Recherche par date de première album
	if isValidDateFormat(query) {
		yearResults := SearchArtistsByFirstAlbumDate(query)
		results = append(results, yearResults...)
	}

	// Recherche par localisations des concerts
	for _, artist := range structs.Artists {
		for location, _ := range artist.Relations.DatesLocations {
			city, country := structs.GetFormattedLocationName(location)
			city = strings.Replace(city, " ", "", -1)
			country = strings.Replace(country, " ", "", -1)
			if strings.Contains(strings.ToLower(city), strings.ToLower(query)) || strings.Contains(strings.ToLower(country), strings.ToLower(query)) {
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
			}
		}

		for _, loc := range artist.Locations {
			if strings.Contains(strings.ToLower(loc), strings.ToLower(query)) {
				matches["Locations"] = true
			}
		}

		if strings.Contains(strings.ToLower(artist.FirstAlbum), strings.ToLower(query)) {
			matches["First Album"] = true
		}

		if strings.Contains(strings.ToLower(fmt.Sprint(artist.CreationDate)), strings.ToLower(query)) {
			matches["Creation Date"] = true
		}

		for location, _ := range artist.Relations.DatesLocations {
			if strings.Contains(strings.ToLower(location), strings.ToLower(query)) {
				matches["Concert Dates"] = true
			}
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

func isValidDateFormat(date string) bool {
	pattern := `^\d{2}/\d{2}/\d{4}$|^\d{2}-\d{2}-\d{4}$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(date)
}

func getYearFromDate(date string) int {
	parts := strings.Split(date, "-")
	yearStr := parts[len(parts)-1]
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		fmt.Println("Error converting year to integer:", err)
		return 0
	}

	return year
}
