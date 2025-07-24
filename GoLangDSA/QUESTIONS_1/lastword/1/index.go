package main

import (
	"fmt"
)

func main() {
	fmt.Print(LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Print(LastWord(" lorem,ipsum "))
	fmt.Print(LastWord(" "))
}

func LastWord(s string) string{

	// Create two pointers starts and ends, return s[starts: ends + 1]

	i := len(s) - 1
	
	for i >= 0 && s[i] == ' ' {
		i--
	}

	ends := i

	// -------------

	for i >= 0 && s[i] != ' ' {
		i--
	}

	starts := i

	return s[starts + 1: ends + 1] + "\n"

}

