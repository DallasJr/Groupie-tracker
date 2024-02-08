package locations

import (
	"groupie-tracker/structs/dates"
)

type Locations struct {
	id        int
	locations []string
	dates     dates.Dates
}
