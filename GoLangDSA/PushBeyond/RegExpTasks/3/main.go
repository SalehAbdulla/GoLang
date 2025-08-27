package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	text, ok := ReadFile("sample.txt")
	if !ok {fmt.Println("Error reading file"); return}

	regExp, err := regexp.Compile(`(?i)hello`)
	if err != nil {fmt.Println("Error compiling file", err); return}

	matches := regExp.FindAllString(text, -1)
	for i, match := range matches {
		fmt.Printf("%d: %s\n", i+1, match)
	}

}

func ReadFile(fileName string) (string, bool) {
	file, err := os.Open(fileName)
	if err != nil {
		return "Error opening file", false
	}
	data := make([]byte, 2048)
	n, err := file.Read(data)
	if err != nil {
		return "Error opening file", false
	}
	return string(data[:n]), true
}
