package main

import (
	"fmt"
)

func main(){
	s := "racecara"
	t := "carracet"
	solution := validAnagram(s, t)
	fmt.Println(solution)
}

func validAnagram(s string, t string) bool {
	// if len not the same, retrun false
	if (len(s) != len(t)) {
		return false
	}

	hashTable := [26]int{}
	for i := range []rune(s) {
		hashTable[s[i] - 'a']++
		hashTable[t[i] - 'a']--
	}

	// O(n) 
	for _, num := range hashTable {
		if num != 0 {
			return false
		}
	}

	return true

}
