package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	readFile, err := os.ReadFile("sample.txt")
	if err != nil {fmt.Println("Error Reading file" + err.Error()); return}
	text := string(readFile)
	
	regExp, err := regexp.Compile(`(?im)(\d+)\s+(\w+)\r?`)
	if err != nil {fmt.Println("Error Compiling" + err.Error()); return}
	
	matchs := regExp.FindAllString(text, -1)
	fmt.Println(matchs)
}
