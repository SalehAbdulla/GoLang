package main

import (
	"fmt"
)

func main(){
	groupAnagrams([]string{"act","pots","tops","cat","stop","hat"})

}

func groupAnagrams(strs []string) [][]string {

	// 1) HashMap | hashTable -> []string
	hashMap := make(map[[26]int][]string)

	for _, word := range strs {
		hashMap[getFrequenct(word)] = append(hashMap[getFrequenct(word)], word)
	}
	// fmt.Println(hashMap)


	// 2) convert them into [][]string
	var result [][]string
	for _, value := range hashMap {
		result = append(result, value)
	}
	fmt.Println(result)


	return result
}

func getFrequenct(s string) [26]int {
	sToRune := []rune(s)

	result := [26]int{}
	for _, char := range sToRune {
		result[char - 'a']++
	}

	return result
}


