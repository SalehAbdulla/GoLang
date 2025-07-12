package main

import (
	"fmt"
)

func main(){
	strs := []string {"act","pots","tops","cat","stop","hat"}
	solution := groupAnagrams(strs)
	fmt.Println(solution)
}

func groupAnagrams(strs []string) [][]string {
	// hashMap | freq [26]int -> ["", "", ""]string
	hashMap := make(map[[26]int][]string)

	for _, word := range strs {
		hashMap[getHashTable(word)] = append(hashMap[getHashTable(word)], word)
	} 

	
	// convert it into Slice
	var result [][]string
	for _, value := range hashMap{
		result = append(result, value)
	}
	
	// fmt.Println(result)

	return result
}

func getHashTable(word string) [26]int {
	result := [26]int{}

	for _, char := range []rune(word){
		result[char - 'a']++
	}

	return result
}