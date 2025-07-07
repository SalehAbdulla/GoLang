package main

import (
	"fmt"
)

func main(){
	strs := []string {"act","pots","tops","cat","stop","hat"};
	fmt.Println(groupAnagram(strs))
}

func groupAnagram(str []string) [][]string {
	hashMap := make(map[[26]int][]string)
	for _, word := range str {
		getCurhashMapKey := getFrequency(word)
		hashMap[getCurhashMapKey] = append(hashMap[getCurhashMapKey], word)
	}

	var result [][]string
	for _, element := range hashMap {
		result = append(result, element)
	}

	return result
}

func getFrequency(str string) [26]int {
	hashTable := [26]int {}
	for i := range str {
		hashTable[str[i] - 'a']++
	}
	return hashTable
}
