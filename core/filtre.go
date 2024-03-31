package core

import (
	"groupie-tracker/structs"
)

var CreationDateRange [2]int
var FirstAlbumDateRange [2]int
var NumberOfMembersRange [2]int
var LocationsCountry []string

func InitializeFiltersValues(artist structs.Artist) {
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
}
