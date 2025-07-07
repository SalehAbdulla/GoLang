package main

import (
	"fmt"
	
)

func main() {
	strs := []string {"act","pots","tops","cat","stop","hat"}
	solution := groupAnagrams(strs)
	fmt.Println(solution)
}

func groupAnagrams(strs []string) [][]string {
	// hashMap [26]int -> []string
	hashMap := make(map[[26]int][]string)

	for _, str := range strs {
		hashMap[getFrequency(str)] = append(hashMap[getFrequency(str)], str)
	}
	
	// take the hashMap's values, append it to [][]string
	var result [][]string
	for _, value := range hashMap {
		result = append(result, value)
	}

	return result

}

func getFrequency(s string) [26]int {
	result := [26]int{}

	for _, char := range []rune(s) {
		result[char - 'a']++
	}

	return result
}
