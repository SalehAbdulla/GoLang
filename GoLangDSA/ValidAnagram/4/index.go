package main


import (
	"fmt"
	"sort"
)

func main() {

}


func isAnagram(s string, t string) bool {
	sortedS := sortString(s)
	sortedT := sortString(t)
	return sortedS == sortedT
}

func sortString(s string) string {
	sToRune := []rune(s)

	sort.Slice(sToRune, func(i, j int) bool {
		return sToRune[i] < sToRune[j]
	})

	return string(sToRune)
}