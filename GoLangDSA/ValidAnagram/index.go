package main 

import (
	"fmt"
	"sort"
)


func main() {
	s := "racecar"
	t := "carrace"
	fmt.Println(validAnagram(s, t))
}


func validAnagram(s string, t string) bool {
	var sToRune []rune
	var tToRune []rune

	sToRune = []rune(s)
	tToRune = []rune(t)

	sort.Slice(sToRune, func(i, j int) bool {
		return sToRune[i] < sToRune[j]
	})

	sort.Slice(tToRune, func(i, j int) bool {
		return tToRune[i] < tToRune[j]
	})

	for i := range sToRune {
		if (sToRune[i] != tToRune[i]) {
			return false
		}
	}

	return true;

}
