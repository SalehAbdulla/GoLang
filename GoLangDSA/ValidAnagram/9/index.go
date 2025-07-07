package main

import (
	"fmt"
)

func main(){
	s := "racecar"
	t := "carrace"
	fmt.Println(validAnagram(s, t))
}

func validAnagram(s string, t string) bool {
	
	if (len(s) != len(t)) {
		return false
	}

	hashTable := [26]int{}

	for i := range []rune(s) {
		hashTable[s[i] - 'a']++
		hashTable[t[i] - 'a']--
	}


	for _, num := range hashTable{
		if (num != 0){
			return false
		}
	}

	return true

}

// Given two strings s and t, return true if the two strings are anagrams of each other, otherwise return false.

// An anagram is a string that contains the exact same characters as another string, but the order of the characters can be different.

// Example 1:

// Input: s = "racecar", t = "carrace"

// Output: true