package main

import (
	"fmt"
)

func main() {
	fmt.Println(RetainFirstHalf("This is the 1st halfThis is the 2nd half"))
	fmt.Println(RetainFirstHalf("A"))
	fmt.Println(RetainFirstHalf(""))
	fmt.Println(RetainFirstHalf("Hello World"))
}

//Instructions
// Write a function called RetainFirstHalf() that takes a string as an argument and returns the first half of this string.

// If the length of the string is odd, round it down.
// If the string is empty, return an empty string.
// If the string length is equal to one, return the string.

func RetainFirstHalf(str string) string {
	if str == "" || len(str) == 1 {
		return str
	}

	return str[:len(str)/2]

}

//$ go run . | cat -e
// This is the 1st half$
// A$
// $
// Hello$
