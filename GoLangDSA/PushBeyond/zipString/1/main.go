package piscine

import (
	"strconv"
)

func ZipString(s string) string {
	var result string
	counter := 1
	for i, char := range s {
		if i+1 < len(s) && char != rune(s[i+1]) {
			result += strconv.Itoa(counter) + string(char)
			counter = 1
		} else if i+1 < len(s) && char == rune(s[i+1]) {
			counter++
		} else if i == len(s)-1 && counter > 0 {
			result += strconv.Itoa(counter) + string(char)
		}
	}
	return result
}
