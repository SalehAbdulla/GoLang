package main

import (
	"github.com/01-edu/z01"
	"os"
)

func getWords(s string) (words []string) {
	var word string
	for _, char := range s {
		if char != ' ' {
			word += string(char)
		} else if word != "" {
			words = append(words, word)
			word = ""
		}
	}

	if word != "" {
		words = append(words, word)
	}

	return
}

func Print(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 || args[0] == "" {
		Print("\n")
		return
	}
	getResult := getWords(args[0])
	if len(getResult) == 1 {
		Print(getResult[0])
		Print("\n")
		return
	}
	firstWord := getResult[0]
	getResult = getResult[1:]

	for _, str := range getResult {
		Print(str)
		Print(" ")

	}
	Print(firstWord)
	Print("\n")

}
