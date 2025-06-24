package main

import (
	"fmt"
)


func main() {
	solution :=  groupAnagram([]string {"act","pots","tops","cat","stop","hat"})
	fmt.Println(solution)
}


func groupAnagram(words []string) [][]string {
	hashMap := make(map[[26]int][]string)

	for _, word := range words {

		var count [26]int
		for _, char := range word {
			count[char - 'a']++
		}
		hashMap[count] = append(hashMap[count], word)
	}

	var result [][]string

	for _, group := range hashMap {
		result = append(result, group)
	}

	return result

}


