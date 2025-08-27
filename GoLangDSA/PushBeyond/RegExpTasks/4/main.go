package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	data, err := os.ReadFile("sample.txt")
	if err != nil {fmt.Println("Error reading file:", err); return}
	text := string(data)



	// fmt.Println(text)

	re, err := regexp.Compile(`(?im)world\r?$`) // \r? because of windows adding \r\n
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}

	matches := re.FindAllString(text, -1)

	for i, match := range matches {
		fmt.Printf("%d: %s\n", i+1, match)
	}
}
