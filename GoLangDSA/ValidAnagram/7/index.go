package main

import (
	"fmt"
)

func main(){

}

func validAnagram(s string, t string) bool {
	if (len(s) != len(t)){
		return false
	}
	
	hashTable := [26]int{};

	for i := 0; i < len(s); i++ {
		hashTable[s[i] - 'a']++;
		hashTable[t[i] - 'a']--;
	}

	for _, element := range hashTable {
		if (element != 0) {
			return false
		}
	}

	return true
}