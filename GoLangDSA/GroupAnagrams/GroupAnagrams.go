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

	hashMap := make(map[string][]string)

	for _, word := range strs {
		sortedWord := sortString(word)
		hashMap[sortedWord] = append(hashMap[sortedWord], word)
	}

	var result [][]string
	for _, element := range hashMap {
		result = append(result, element)
	}

	return result
}


func sortString(word string) string {

	wordToRune := []rune(word)

	sort.Slice(wordToRune, func(i, j int) bool {
		return wordToRune[i] < wordToRune[j]
	})
	
	return string(wordToRune)
}