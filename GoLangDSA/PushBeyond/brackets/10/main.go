package main

import (
	"fmt"
	"os"
)

func BracketBalanced(text string) bool {
	var stack []rune
	for _, char := range text {
		if char == '[' || char == '(' || char == '{' {
			stack = append(stack, char)
		} else if char == ']' {
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1] // pop
		} else if char == ')' {
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		} else if char == '}' {
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println()
		return
	}
	for _, arg := range args {
		if BracketBalanced(arg) {
			fmt.Println("OK")
		} else {
			fmt.Println("Error")
		}
	}
}
