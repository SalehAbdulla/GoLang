// Using Sort Algorithm

package main


import (
	"fmt"
	"sort"
)


func main() {
	strs := []string {"act","pots","tops","cat","stop","hat"}
	fmt.Println(groupAnagrams(strs))
}

func groupAnagrams(strs []string) [][]string {

	hashMap := make(map[string][]string)

	for _, word := range strs {
		sortedWord := sortString(word)
		hashMap[sortedWord] = append(hashMap[sortedWord], word)
	}

	var result [][]string
	for _, group := range hashMap {
		result = append(result, group)
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