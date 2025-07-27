package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main(){
	args := os.Args[1:]
	
	// If the number of arguments is different from 2, the program displays nothing.
	if len(args) != 2 {return}

	word1 := args[0]
	word2 := args[1]


	i := 0 // flag
	for _, char := range word2 {
							// if we encountered char that same as char exists in word2
							// it trakes a little time until you get it
							// its a strategy to memorise, add it to your Brain :)
		if i < len(word1) && char == rune(word1[i]) {i++}
	}

	if i == len(word1) {
		for _, char := range word1 {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}



}










//Instructions

// Write a program that takes two string and checks whether it is possible to write the first string with characters from the second string. This rewrite must respect the order in which these characters appear in the second string.

// If it is possible, the program displays the string followed by a newline ('\n'), otherwise it simply displays nothing.



// Usage
// $ go run . 123 123
// 123
// $ go run . faya fgvvfdxcacpolhyghbreda
// faya
