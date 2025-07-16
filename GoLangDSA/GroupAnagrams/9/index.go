package main

import (
	"fmt"
)

func main(){
	strs := []string {"act","pots","tops","cat","stop","hat"}
	solution := groupAnagrams(strs)
	fmt.Println(solution)
}

//Input: strs = ["act","pots","tops","cat","stop","hat"]
// Output: [["hat"],["act", "cat"],["stop", "pots", "tops"]]

func groupAnagrams(str []string)[][]string {
	hashMap := make(map[[26]int][]string)
	for _, word := range str {
		hashMap[getHashTable(word)] = append(hashMap[getHashTable(word)], word)
	}

	var result [][]string
	for _, value := range hashMap {
		result = append(result, value)
	}

	return result
}

func getHashTable(str string)[26]int {
	hashTable := [26]int{}
	for _, char := range str {
		hashTable[char - 'a']++
	}
	return hashTable
}



