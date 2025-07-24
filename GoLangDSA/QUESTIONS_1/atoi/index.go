package main

import (
	"fmt"
)

func main() {
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))
}

func Atoi(s string) int {
	answer := 0
	isNegative := false

	if s[0] == '-' {
		isNegative = true
	}
	
	if s[0] == '-' && s[1] == '-' || s[0] == '+' && s[1] == '+' {
		return 0
	}

	for _, char := range s {

		if char == '-' || char == '+' {
			continue
		
		} else if (char >= '0' && char <= '9') {
			answer *= 10
			answer += int(char) - '0'
		} else {
			return 0
		}
	}

	if isNegative {
		return answer * -1
	} else {
		return answer
	}

}

