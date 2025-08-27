package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	text, err := os.ReadFile("sample.txt")
	if err != nil {fmt.Println("Error Reading file"); return}
	data := string(text)

	regExp, err := regexp.Compile(`\(\s*([^\s(),]+)\s*(?:,\s*(\d+))?\s*\)`)
	if err != nil {fmt.Println("Error Compiling reqexp", err.Error()); return}

	matches := regExp.FindAllString(data, -1)
	fmt.Println(matches)
}