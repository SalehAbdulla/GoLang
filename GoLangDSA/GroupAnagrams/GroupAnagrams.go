package main

import (
	"fmt"
	"sort"
)

func main(){
	strs := []string {"act","pots","tops","cat","stop","hat"}
	groupAnagrams(strs)
}


func groupAnagrams(strings []string) [][]string {
	hash := make(map[string][]string)

	for _, word := range strings {
		key := sortWord(word)
		hash[key] = append(hash[key],word)
	}

	var result [][]string;

	

}


func sortWord(word string) string {
	wordToRune := []rune(word)

	// return void
	sort.Slice(wordToRune, func(i, j int) bool {
		return wordToRune[i] < wordToRune[j]
	})

	// converts it from rune intro string
	return string(wordToRune)
}
