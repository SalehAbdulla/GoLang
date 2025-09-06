package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "mohammed (up) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair."

	textToSlice := strings.Fields(text)

	for i, str := range textToSlice {
		if str == "(cap)" {
			befWord := textToSlice[i-1]
			textToSlice[i-1] = strings.ToUpper(befWord[:1]) + strings.ToLower(befWord[1:])
			break
		} else if str == "(up)" {
			befWord := textToSlice[i-1]
			textToSlice[i-1] = strings.ToUpper(befWord) 
		}
	}
	fmt.Println(textToSlice)
}




