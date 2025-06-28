package main

import (
	"fmt"
)

func main() {
	strs := []string {"eat","tea","tan","ate","nat","bat"}
	fmt.Println(groupAnagrams(strs))

}	

func groupAnagrams(strs []string) [][]string {
	hashMap := make(map[[26]int][]string)

	for _, element := range strs {
		elementToFrequency := getFrequency(element);
		hashMap[elementToFrequency] = append(hashMap[elementToFrequency], element)
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
	for _, char := range sToRune {
		count[char - 'a']++
	}

	return count
}