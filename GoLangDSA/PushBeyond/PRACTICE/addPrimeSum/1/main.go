package main

import (
	"os"
	

	"github.com/01-edu/z01"
)

//- If the number of arguments is different from 1, or if the argument is not a positive number,
// //the program displays `0` followed by a newline.



func main(){
	args := os.Args[1:]
	if len(args) != 1 {
		z01.PrintRune('0')
		return
	}

}


func Atoi(number string) int {
	if number == "" {return -1}
	if number == "0" {return 0}
	// ---------
	isNegative := false
	if number[0] == '-' {
		isNegative = true
		number = number[1:]
	}
	// ---------
	result := 0
	for _, char := range number {
		if char >= '0' && char <= '9' {
			result *= 10
			result += int(char - '0')
		} else {
			return -1
		}
	}

	// ----------
	if isNegative {
		return result * -1
	}
	return result 
}


