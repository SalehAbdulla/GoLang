package main

import (
	"github.com/01-edu/z01"
	"os"
)

func Print(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		return
	}

	word1 := args[0]
	word2 := args[1]

	i := 0
	for _, char := range word2 {
		if i < len(word1) && char == rune(word1[i]) {
			i++
		}
	}

	if i == len(word1) {
		Print(word1)
		Print("\n")
	}
}
