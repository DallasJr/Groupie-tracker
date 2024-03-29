package structs

type Relation struct {
	id              int              `json:"id"`
	Dates_Locations []DatesLocations `json:"datesLocations"`
}

type DatesLocations struct {
	location    string   `json:"location"`
	Table_Dates []string `json:"dates"`
}
