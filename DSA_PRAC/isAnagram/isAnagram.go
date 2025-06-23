package main

import (
	"fmt"
	"sort"
)

func main() {
	s := "racecar"
	t := "carrace"
	var solution bool = isAnagram(s, t)
	fmt.Println(solution)
}

func isAnagram(s string, t string) bool {

	if (len(s) != len(t)) {
		return false
	}

	sToRune := []rune(s)
	tToRune := []rune(t)

	sort.Slice(sToRune, func(i, j int) bool {
		return sToRune[i] < sToRune[j]
	})

	sort.Slice(tToRune, func(i, j int) bool {
		return tToRune[i] < tToRune[j];
	})

	for i := range sToRune {
		if (sToRune[i] != tToRune[i]) {
			return false
		}
	}

	return true

}


// Input: s = "racecar", t = "carrace"
// Output: true