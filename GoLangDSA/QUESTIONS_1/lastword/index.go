package main

import (
	"fmt"
)

func main() {
	fmt.Print(LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Print(LastWord(" lorem,ipsum "))
	fmt.Print(LastWord(" "))
}

// LastWord returns the last word in the string followed by a newline
func LastWord(s string) string {

	// Get the end of the word
	i := len(s) - 1
	for i >= 0 && s[i] == ' ' {
		i--
	}
	ends := i // ends saved

	// -----------------

	for i >= 0 && s[i] != ' ' {
		i--
	}
	starts := i + 1

	return s[starts:ends + 1] + "\n"

}
