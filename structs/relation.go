package structs

type Relation struct {
	id             int
	datesLocations []DatesLocations
}

type DatesLocations struct {
	location string
	dates    []string
}
