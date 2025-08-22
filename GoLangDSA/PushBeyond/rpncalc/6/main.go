package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RpnCalc(strs []string) (result int, ok bool) {
	var stack []int
	for _, str := range strs {
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
	} else {
		return -1, false
	}
}

func removeSpaces(strs []string) (result []string) {
	for _, str := range strs {
		if str == "" || str == " " {
			continue
		} else {
			result = append(result, str)
		}
	}
	return
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Error")
		return
	}
	operation := strings.Split(args[0], " ")
	operation = removeSpaces(operation)
	calc, ok := RpnCalc(operation)
	if !ok {
		fmt.Println("Error")
	} else {
		fmt.Println(calc)
	}

}
