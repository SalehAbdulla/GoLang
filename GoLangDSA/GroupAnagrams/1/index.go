package main

import (
	"fmt"
	"sort"
)

func main() {
	
	strs := []string {"act","pots","tops","cat","stop","hat"}
	solution := groupAnagrams(strs)	

	fmt.Println(solution)
}

func groupAnagrams(strs []string) [][]string {

	// Create HashMap
	hashmap := make(map[string][]string)

	// Loop through each word
	for _, element := range strs {
		// Sort each word - sorted word as a key in the hashmap 
		sortedWord := sortString(element)

		// loopback again, if catch same word again, append it - orginal word as a value
		hashmap[sortedWord] = append(hashmap[sortedWord], element);
	}

	// Wrap it into new slice -- must loop, we can't add the hashmap directly in Go, then return it back

	var result [][]string
	for _, group := range hashmap{
		result = append(result, group)
	}
	return result;
}


func sortString(word string) string {
	wordToRune := []rune(word)

	sort.Slice(wordToRune, func(a, b int) bool {
		return wordToRune[a] < wordToRune[b]
	})

	return string(wordToRune)
}