package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	text, err := os.ReadFile("sample.txt")
	if err != nil {
		fmt.Println("Error reading file")
		return
	}
	data := string(text)

	// ----------

	// (    cap)
	// (low,3)
	// (up, 10)
	// (trim-text, 2)

	regExp, err := regexp.Compile(`\(\s*([^\s(),]+)\s*(?:,\s*(\d+))?\s*\)`)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	matches := regExp.FindAllString(data, -1)

	for i, match := range matches {
		fmt.Printf("%d: %s \n", i+1, match)
	}

}
