package main

import (
	"encoding/json"
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
	artists = []Artist{
		{
			ID:           1,
			Name:         "Queen",
			Members:      []string{"Freddie Mercury", "Brian May", "John Daecon", "Roger Meddows-Taylor", "Mike Grose", "Barry Mitchell", "Doug Fogie"},
			CreationDate: 1970,
			FirstAlbum:   "14-12-1973",
			Locations:    "https://groupietrackers.herokuapp.com/api/locations/1",
			ConcertDates: "https://groupietrackers.herokuapp.com/api/dates/1",
			Relations:    "https://groupietrackers.herokuapp.com/api/relation/1",
		},
	}

	http.HandleFunc("/search", searchHandler)
	http.ListenAndServe(":8080", nil)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")

	if query == "" {
		http.Error(w, "Le terme de recherche est vide", http.StatusBadRequest)
		return
	}

	results := searchArtistsByName(query)

	response, err := json.Marshal(results)
	if err != nil {
		http.Error(w, "Erreur lors de la conversion en JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func searchArtistsByName(query string) []Artist {
	var results []Artist
	for _, artist := range artists {
		if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(query)) {
			results = append(results, artist)
		}
	}
	return results
}
