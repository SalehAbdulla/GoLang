package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	}
	input := args[0]
	if input == " " {
		return
	}
	inputToInt, err := strconv.Atoi(input)

	if err != nil {
		return
	}
	if inputToInt < 2 {
		return
	}
	var nums []int

	i := 2
	for i*i <= inputToInt {
		for inputToInt%i == 0 {
			nums = append(nums, i)
			inputToInt /= i
		}
		i++
	}

	if inputToInt > 2 {
		nums = append(nums, inputToInt)
	}

	for i, n := range nums {
		fmt.Print(n)
		if i != len(nums)-1 {
			fmt.Print("*")
		}
	}
	fmt.Println()
}
