package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {return}

	runes := []rune(args[0])
	words := []string{}
	word := ""

	// Build words from input
	for _, r := range runes {
		if r != ' ' && r != '\t' {
			word += string(r)
		} else if word != "" {
			words = append(words, word)
			word = "" // reset word
		}
	}
    
	// Append last word if needed
	if word != "" {
		words = append(words, word)
	}

	// Print words with 3 spaces in between
	for i, w := range words {
		for _, r := range w {
			z01.PrintRune(r)
		}
		if i != len(words)-1 {
			z01.PrintRune(' ')
			z01.PrintRune(' ')
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
