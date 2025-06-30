package main

import (
	"fmt"
	"sort"
)

func main() {
	s := "racecar";
	t := "carrace";
	solution := validAnagram(s, t);
	fmt.Println(solution);
}

func validAnagram(s string, t string) bool {
	sToFrequency := getFrequency(s)
	tToFrequency := getFrequency(t)
	return sToFrequency == tToFrequency;
}

func getFrequency(word string) string {
	wordToRune := []rune(word)

	sort.Slice(wordToRune, func(i, j int) bool {
		return wordToRune[i] < wordToRune[j]
	})

	return string(wordToRune)

}
