// Solution with sorting


package main


import (
	"fmt"
	"sort"
)


func main2() {
	s := "jar"
	t := "jam"

	fmt.Println(validAnagram(s, t))
}

func validAnagram(s, t string) bool {
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