package structs

import "strings"

type Locations struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
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
