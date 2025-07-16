package main

import (
	"fmt"
)

func main(){
	s := "racecar"
	t := "carracf"
	fmt.Println(validAnagram(s, t))
}

func validAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	hashTable := [26]int{}

	for i, _ := range s {
		hashTable[s[i] - 'a']++
		hashTable[t[i] - 'a']--
	}
	
	for _, num := range hashTable {
		if num != 0 {
			return false
		}
	}
	return true
}

//Input: s = "racecar", t = "carrace"
// Output: true
