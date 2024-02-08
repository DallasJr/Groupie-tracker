package locations

import "groupie-tracker/dates"

type Locations struct {
	id        int
	locations []string
	dates     dates.Dates
}
