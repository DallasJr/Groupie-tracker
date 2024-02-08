package structs

type Artists struct {
	id           int
	image        string
	name         string
	members      []string
	creationDate int
	firstAlbum   string
	locations    Locations
	concertDates Dates
	relations    Relation
}
