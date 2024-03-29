package structs

import "strings"

type Relations struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

func GetFormatted(s string) string {
	s = strings.Replace(s, "-", " ", -1)
	s = strings.Replace(s, "_", " ", -1)
	words := strings.Fields(s)
	for i, word := range words {
		words[i] = strings.ToUpper(string(word[0])) + word[1:]
	}
	return strings.Join(words, " ")
}
