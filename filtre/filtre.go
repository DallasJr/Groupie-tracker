package filtre

import (
	"groupie-tracker/structs"
	"sort"
	"time"
)

type Filters struct {
	CreationDateMin float64
	CreationDateMax float64
	FirstAlbumMin   float64
	FirstAlbumMax   float64
	NumMembers      int
	Locations       []string
}

func parseFirstAlbumDate(dateStr string) float64 {
	t, err := time.Parse("02-01-2006", dateStr)
	if err != nil {
		return 0
	}
	return float64(t.Year())
}

func FilterArtists(artists []structs.Artist, filters Filters) []structs.Artist {
	filtered := make([]structs.Artist, 0)

	for _, artist := range artists {
		firstAlbumYear := parseFirstAlbumDate(artist.FirstAlbum)

		if float64(artist.CreationDate) >= filters.CreationDateMin &&
			(float64(artist.CreationDate) <= filters.CreationDateMax || filters.CreationDateMax == 0) &&
			firstAlbumYear >= filters.FirstAlbumMin &&
			(firstAlbumYear <= filters.FirstAlbumMax || filters.FirstAlbumMax == 0) &&
			(len(artist.Members) == filters.NumMembers || filters.NumMembers == 0) &&
			(len(filters.Locations) == 0 || intersects(filters.Locations, artist.Locations.Locations)) {
			filtered = append(filtered, artist)
		}
	}

	SortArtistsByCreationDate(filtered)
	return filtered
}

func intersects(a, b []string) bool {
	for _, v := range a {
		for _, v2 := range b {
			if v == v2 {
				return true
			}
		}
	}
	return false
}

func SortArtistsByCreationDate(artists []structs.Artist) {
	sort.Slice(artists, func(i, j int) bool {
		return artists[i].CreationDate < artists[j].CreationDate
	})
}
