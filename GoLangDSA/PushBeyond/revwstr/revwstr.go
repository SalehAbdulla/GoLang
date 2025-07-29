package main

// If the number of arguments is different from 1, the program will display nothing.
import (
	"os"
	"github.com/01-edu/z01"
)

func main(){
	args := os.Args[1:]
	if len(args) != 1 {return}
	input := args[0]

	words := []string{}
	word := ""

	for _, char := range input {
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

	result := ""
	for i := len(words) - 1; i >= 0; i-- {
		result += string(words[i])
		if i != 0 {
			result += string(" ")
		}
	}

	for _, char := range result {
		z01.PrintRune(char)
	}

	z01.PrintRune('\n')

}

