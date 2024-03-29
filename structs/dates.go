package structs

import "fmt"

type Dates struct {
	id    int      `json:"id"`
	Dates []string `json:"dates"`
}

func CheckLoadedDates() {
	var dates []Dates
	if len(dates) == 0 {
		fmt.Println("Aucune date n'a été chargée.")
		return
	}

	fmt.Println("Dates chargées :")
	for _, date := range dates {
		fmt.Printf("- ID: %d, Dates: %v\n", date.id, date.Dates)
	}
}
