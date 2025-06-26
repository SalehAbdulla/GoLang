package main


import (
	"fmt"
)


func main() {
	strs := []string {"act","pots","tops","cat","stop","hat"}
	fmt.Println(groupAnagrams(strs))
}

func groupAnagrams(strs []string) [][]string {
	hashMap := make(map[[26]int][]string)

	for _, word := range strs {
		wordToFrequency := getFrequency(word)
		hashMap[wordToFrequency] = append(hashMap[wordToFrequency], word)
	}

	var result [][]string
	for _, group := range hashMap {
		result = append(result, group)
	}
	return result

}


func getFrequency(s string) [26]int {
	sToRune := []rune(s)
	var count [26]int

	for _, element := range sToRune {
		count[element - 'a']++
	}

	return count
	
}