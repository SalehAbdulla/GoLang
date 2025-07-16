package main

import (
	"fmt"
)

func main() {
	fmt.Println(CountChar("Hello World", 'l'))
	fmt.Println(CountChar("5  balloons", 5))
	fmt.Println(CountChar("   ", ' '))
	fmt.Println(CountChar("The 7 deadly sins", '7'))
}

func CountChar(str string, char rune) int {

	//if the character is not in the string return 0
	isExist := false
	for _, c := range str {
		if c == char {
			isExist = true
			break
		} 
	}
	if isExist == false {
		return 0
	}
	// if the string is empty return 0
	if str == "" {
		return 0
	}
	hashMap := make(map[rune]int)
	for _, char := range str {
		hashMap[char]++
	}
	return hashMap[char]
}
