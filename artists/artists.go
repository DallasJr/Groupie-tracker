package artists

import (
	"groupie-tracker/dates"
	"groupie-tracker/locations"
	"groupie-tracker/relation"
)

type Artists struct {
	id           int
	image        string
	name         string
	members      []string
	creationDate int
	firstAlbum   string
	locations    locations.Locations
	concertDates dates.Dates
	relations    relation.Relation
}
