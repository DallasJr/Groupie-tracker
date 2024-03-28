package structs

type Relation struct {
	id               int      `json:"id"`
	Relationlocation []string `json:"location"`
	Table_Dates      []string `json:"dates"`
}
