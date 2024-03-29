package structs

type Relation struct {
	id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}
