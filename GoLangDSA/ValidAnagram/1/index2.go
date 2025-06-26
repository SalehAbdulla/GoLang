// Solution using Frequency

package main


import (
	"fmt"
)



func main() {
	s := "racecar"
	t := "carrace"
	solution := validAnagram2(s, t)
	fmt.Println(solution)
}


func validAnagram2(s, t string) bool {

	if (len(s) != len(t)) {
		return false;
	}

	count := [26]int{};

	for _, char := range s {
		count[char - 'a']++
	}

	for _, char := range t {
		count[char - 'a']--
	}

	for _, element := range count {
		if (element != 0) {
			return false
		}
	}

	return true

}	