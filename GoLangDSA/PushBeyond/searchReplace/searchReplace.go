package main

//$ go run . "hella there" "a" "o"
// hello there

// Write a program that takes 3 arguments, the first argument is a string in which a letter

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	getArgs := os.Args[1:] // get rid of location

	//// If the number of arguments is different from 3, the program displays nothing.
	if len(getArgs) != 3 {
		z01.PrintRune(' ')
		return
	}

	word := getArgs[0]        // Actual word
	wrongChar := getArgs[1]   // takes wrongChar
	correctChar := getArgs[2] // takes correctChar

	// --------------------
	//// If the second argument is not contained in the first one (the string)
	// Then the program rewrites the string followed by a newline ('\n').
	isExist := false
	for _, char := range word {
		if string(char) == wrongChar {
			isExist = true
		}
	}

	if !isExist {
		for _, c := range word {
			z01.PrintRune(c)
		}
		z01.PrintRune('\n')
		return
	}
	// --------------------
	////(the 2nd argument) will be replaced by another one (the 3rd argument).
	result := []rune(word)
	for i, char := range result {
		if string(char) == wrongChar {
			// string to rune
			wrongCharToRune := []rune(correctChar)
			result[i] = wrongCharToRune[0]
		}
	}
	// Print Result
	for i := 0; i < 1; i++ {
		for _, c := range result {
			z01.PrintRune(c)
		}
        z01.PrintRune('\n')
	}

}
