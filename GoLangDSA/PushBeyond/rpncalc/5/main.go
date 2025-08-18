package main

//$ go run . "1 2 * 3 * 4 +" | cat -e
// 10$

import (
	"fmt"
	"os"
	"strconv"
)

func CalcRpn(text []string) (int, bool) {
	var stack []int
	for _, str := range text {
		if str == " " {
			continue
		}
		strToInt, err := strconv.Atoi(str)
		if err == nil {
			stack = append(stack, strToInt)
			continue
		}
		if len(stack) < 2 {
			return -1, false
		}
		switch str {
		case "+":
			stack[len(stack)-2] += stack[len(stack)-1]
			stack = stack[:len(stack)-1] // pop
		case "-":
			stack[len(stack)-2] -= stack[len(stack)-1]
			stack = stack[:len(stack)-1] // pop
		case "*":
			stack[len(stack)-2] *= stack[len(stack)-1]
			stack = stack[:len(stack)-1] // pop
		case "/":
			stack[len(stack)-2] /= stack[len(stack)-1]
			stack = stack[:len(stack)-1] // pop
		case "%":
			stack[len(stack)-2] %= stack[len(stack)-1]
			stack = stack[:len(stack)-1] // pop
		default:
			return -1, false
		}
	}
	if len(stack) == 1 {
		return stack[0], true
	}
	return -1, false
}

func removeSpaces(s string) (result []string) {
	var word string
	for _, char := range s {
		if char != ' ' {
			word += string(char)
		} else if word != "" {
			result = append(result, word)
			word = ""
		}
	}
	if word != "" {
		result = append(result, word)
	}
	return
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Error")
		return
	}

	solution, ok := CalcRpn(removeSpaces(args[0]))
	if !ok {
		fmt.Println("Error")
		return
	}
	fmt.Println(solution)
}
