package main

import (
	"fmt"
)

func main() {
	fmt.Println(CheckNumber("Hello"))
	fmt.Println(CheckNumber("Hello1"))
}

// Write a function that takes a string as an argument and returns true if the string contains any number, otherwise return false.

func CheckNumber(str string) bool {

	isExist := false
	for _, char := range []rune(str) {
		if char >= '0' && char <= '9' {
			isExist = true
			break
		} 
	}

	return isExist

}


