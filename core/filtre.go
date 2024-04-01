package core

import (
	"groupie-tracker/structs"
	"strings"
)

var CreationDateRange [2]int
var CreationDateValue [2]float64

var FirstAlbumDateRange [2]int
var FirstAlbumDateValue [2]float64

var NumberOfMembersRange [2]int
var NumberOfMembersValue [2]float64

var LocationsCountry []string
var LocationsCountryChecked []string

func InitializeFiltersValues(artist structs.Artist, loop bool) {
	// Récuperer les valeurs min et max
	if loop {
		if CreationDateRange[0] == 0 {
			CreationDateRange[0] = artist.CreationDate
			CreationDateRange[1] = artist.CreationDate
		} else {
			if CreationDateRange[0] > artist.CreationDate {
				CreationDateRange[0] = artist.CreationDate
			} else if CreationDateRange[1] < artist.CreationDate {
				CreationDateRange[1] = artist.CreationDate
			}
		}
		if FirstAlbumDateRange[0] == 0 {
			FirstAlbumDateRange[0] = getYearFromDate(artist.FirstAlbum)
			FirstAlbumDateRange[1] = getYearFromDate(artist.FirstAlbum)
		} else {
			if FirstAlbumDateRange[0] > getYearFromDate(artist.FirstAlbum) {
				FirstAlbumDateRange[0] = getYearFromDate(artist.FirstAlbum)
			} else if FirstAlbumDateRange[1] < getYearFromDate(artist.FirstAlbum) {
				FirstAlbumDateRange[1] = getYearFromDate(artist.FirstAlbum)
			}
		}
		if NumberOfMembersRange[0] == 0 {
			NumberOfMembersRange[0] = len(artist.Members)
			NumberOfMembersRange[1] = len(artist.Members)
		} else {
			if NumberOfMembersRange[0] > len(artist.Members) {
				NumberOfMembersRange[0] = len(artist.Members)
			} else if NumberOfMembersRange[1] < len(artist.Members) {
				NumberOfMembersRange[1] = len(artist.Members)
			}
		}
	} else {

		// Défini les valeurs par défaut
		CreationDateValue[0] = float64(CreationDateRange[0])
		CreationDateValue[1] = float64(CreationDateRange[1])
		FirstAlbumDateValue[0] = float64(FirstAlbumDateRange[0])
		FirstAlbumDateValue[1] = float64(FirstAlbumDateRange[1])
		NumberOfMembersValue[0] = float64(NumberOfMembersRange[0])
		NumberOfMembersValue[1] = float64(NumberOfMembersRange[1])
	}
}

func GetFiltered(artists []structs.Artist) []structs.Artist {
	var filteredArtists []structs.Artist
	for _, artist := range artists {
		creationDate := float64(artist.CreationDate)
		firstAlbumDate := float64(getYearFromDate(artist.FirstAlbum))
		numberOfMembers := float64(len(artist.Members))
		locations := artist.Locations
		if creationDate >= CreationDateValue[0] &&
			creationDate <= CreationDateValue[1] &&
			firstAlbumDate >= FirstAlbumDateValue[0] &&
			firstAlbumDate <= FirstAlbumDateValue[1] &&
			numberOfMembers >= NumberOfMembersValue[0] &&
			numberOfMembers <= NumberOfMembersValue[1] {
			if len(LocationsCountryChecked) != 0 {
				good := false
				for _, location := range locations {
					if good {
						break
					}
					for _, countryChecked := range LocationsCountryChecked {
						if strings.Contains(strings.ToLower(location), strings.ToLower(countryChecked)) {
							good = true
							break
						}
					}
				}
				if good {
					filteredArtists = append(filteredArtists, artist)
				}
			} else {
				filteredArtists = append(filteredArtists, artist)
			}
		}
	}
	return filteredArtists
}
