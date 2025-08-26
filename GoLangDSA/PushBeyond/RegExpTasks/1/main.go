package main

import (
	"fmt"
	"regexp"
)

func main() {
	text := "Ahmed went to Wonderland with Bob and Charlie."
	regExp, err := regexp.Compile(`([A-Z][a-z]*)`)

	if err != nil {
		fmt.Println("Error Compiling regular expression: ", err)
		return
	}
	matches := regExp.FindAllString(text, -1)

	for i, match := range matches {
		fmt.Printf("%d: %s\n", i+1, match)
	}

}
