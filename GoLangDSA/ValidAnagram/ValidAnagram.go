package main

import (
	"slices"
	"fmt"
)

// Input: s = "racecar", t = "carrace"
// Output: true

func main() {
	s := "racecar"
	t := "carrace"
	fmt.Println(isValidAnagram(s, t))
}

func isValidAnagram(s string, t string) bool {
	sToRunes := []rune(s)
	tToRunes := []rune(t)

	if len(sToRunes) != len(tToRunes) {
		return false
	}

	slices.Sort(sToRunes);
	slices.Sort(tToRunes);

	
	for i := range sToRunes {
		if (sToRunes[i] != tToRunes[i]) {
			return false;
		}
	}

	return true;


}
