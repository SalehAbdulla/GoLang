package main

import (
	"fmt"
)

func main() {
	fmt.Println(BasicAtoi("12345"))
	fmt.Println(BasicAtoi("0000000012345"))
	fmt.Println(BasicAtoi("000000"))
}

// Write a function that simulates the behaviour of the Atoi function in Go. Atoi transforms a number defined as a string into a number defined as an int.

func BasicAtoi(s string) int {
	answer := 0

	for _, char := range s {
		// fmt.Println(answer)
		actualNumber := (int(char) - '0') 
		answer *= 10
		answer += actualNumber
	}

	return answer

}

