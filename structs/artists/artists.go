package artists

import (
	"groupie-tracker/structs/dates"
	"groupie-tracker/structs/locations"
	"groupie-tracker/structs/relation"
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
