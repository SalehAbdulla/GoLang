package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args[1:]
	if len(args) != 2 || args[1] == "" {
		return
	}

	regExp := args[0]
	if regExp[0] != '(' || regExp[len(regExp)-1] != ')' {
		return
	}
	regExp = regExp[1 : len(regExp)-1]

	var expToSlice []string

	if Contains(regExp, "|") {
		expToSlice = append(expToSlice, Split(regExp, "|")...)
	} else {
		expToSlice = append(expToSlice, regExp)
	}

	textsToSlice := Split(args[1], " ")

	getResult := GetWordContains(expToSlice, textsToSlice)

	for i, word := range getResult {
		if EndsWithSymbol(word) {
			word = word[:len(word)-1] // pop
			fmt.Printf("%d: %s\n", i+1, word)
		} else {
			fmt.Printf("%d: %s\n", i+1, word)
		}
	}

}

func EndsWithSymbol(w string) bool {
	return !(w[len(w)-1] >= 'a' && w[len(w)-1] <= 'z') && !(w[len(w)-1] >= 'A' && w[len(w)-1] <= 'Z')
}

func Contains(text, toFind string) bool {
	for i := 0; i <= len(text)-len(toFind); i++ {
		if toFind == text[i:i+len(toFind)] {
			return true
		}
	}
	return false
}

func GetWordContains(exps, words []string) (result []string) {
	for _, word := range words {
		for _, exp := range exps {
			if Contains(word, exp) {
				result = append(result, word)
			}
		}
	}
	return
}

func Split(words, splitter string) (result []string) {
	var word string
	for _, char := range words {
		if string(char) != splitter {
			word += string(char)
		} else if word != "" {
			result = append(result, word)
			word = ""
		}
	}
	if word != "" {
		result = append(result, word)
	}
	return
}
