package core

import (
	"fmt"
	"strconv"
	"strings"
)

func ContainsString(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

func getYearFromDate(date string) int {
	parts := strings.Split(date, "-")
	yearStr := parts[len(parts)-1]
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		fmt.Println("Error converting year to integer:", err)
		return 0
	}

	return year
}

func FirstLetterUpper(s string) string {
	words := strings.Fields(s)
	for i, word := range words {
		words[i] = strings.ToUpper(string(word[0])) + word[1:]
	}
	s = strings.Join(words, " ")
	return s
}
