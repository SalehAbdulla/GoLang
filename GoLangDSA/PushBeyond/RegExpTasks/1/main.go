package main

import (
	"fmt"
	"regexp"
)

func main() {
	textToBeModified := "12 Apples, 10 Bananas, and 6 watermelons"
	regExp, err := regexp.Compile(`\d+`) // 0 or more real digits
	if err != nil {
		fmt.Println("Error Compiling regular expression", err)
		return
	}
	matches := regExp.FindAllString(textToBeModified, -1)
	for i, match := range matches {
		fmt.Printf("%d: %s\n", i+1, match)
	}
}
