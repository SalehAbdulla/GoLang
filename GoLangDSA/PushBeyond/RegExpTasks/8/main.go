package main

import (
	"fmt"
	"os"
	"regexp"
)


func main() {
	text, err := os.ReadFile("sample.txt")
	if err != nil {fmt.Println(err.Error()); return}

	regExp, err := regexp.Compile(`(?im)\(\s*([A-Za-z]+)\s*(?:,\s*([0-9]+))?\s*\)`)
	
	if err != nil {fmt.Println(err.Error()); return}
	matches := regExp.FindAllString(string(text), -1)
	for i, match := range matches {
		fmt.Printf("%d: %s\n", i+1, match)
	}


}


