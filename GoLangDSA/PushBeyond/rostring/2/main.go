package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 || args[0] == "" {
		z01.PrintRune('\n')
		return
	}
	var result []string
	word := ""
	for _, char := range args[0] {
		if char != ' ' {
			word += string(char)
		} else if word != "" {
			result = append(result, word)
			word = ""
		}
	}
	if word != "" {
		result = append(result, word)
	}
	firstWord := result[0]
	result = result[1:]
	for _, w := range result {
		for _, c := range w {
			z01.PrintRune(c)
		}
		z01.PrintRune(' ')
	}
	for _, c := range firstWord {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
