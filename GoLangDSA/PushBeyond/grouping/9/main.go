package main

import (
	"fmt"
	"os"
)

func Split(text, splitter string) []string {
	var words []string
	var word string
	for _, char := range text {
		if string(char) != splitter {
			word += string(char)
		} else if word != "" {
			words = append(words, word)
			word = ""
		}
	}
	if word != "" {
		words = append(words, word)
	}
	return words
}

func Contains(text, toFind string) bool {
	for i := 0; i <= len(text)-len(toFind); i++ {
		if toFind == text[i:i+len(toFind)] {
			return true
		}
	}
	return false
}

func EndsWithSymbol(t string) bool {
	return (t[len(t)-1] >= 'a' && t[len(t)-1] <= 'z') || (t[len(t)-1] >= 'A' && t[len(t)-1] <= 'Z')
}

func Search(exps, texts []string) (result []string) {
	for _, text := range texts {
		for _, exp := range exps {
			if Contains(text, exp) {
				result = append(result, text)
			}
		}
	}
	return
}

func PrintAsNeeded(strs []string) {
	for i, str := range strs {
		if !EndsWithSymbol(str) {
			str = str[:len(str)-1]
		}
		fmt.Printf("%d: %s\n", i+1, str)
	}
}

//If the number of arguments is different from 2, if the regular expression is not valid,
//if the last argument is empty or if there are no matches, the program should print nothing.

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		return
	}
	exp := args[0]

	if exp[0] != '(' || exp[len(exp)-1] != ')' {
		return
	}
	exp = exp[1 : len(exp)-1]

	var expToSlice []string
	if Contains(exp, "|") {
		expToSlice = append(expToSlice, Split(exp, "|")...)
	} else {
		expToSlice = append(expToSlice, exp)
	}

	var textToSlice []string
	textToSlice = append(textToSlice, Split(args[1], " ")...)
	PrintAsNeeded(Search(expToSlice, textToSlice))

}
