package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Invalid input.")
		return
	}

	nums := args[0]
	if len(nums) < 3 || nums[0] != '[' || nums[len(nums)-1] != ']' {
		fmt.Println("Invalid input.")
		return
	}

	nums = nums[1 : len(nums)-1] // removes [] brackets
	nums = strings.ReplaceAll(nums, " ", "")

	numsSlice = strings.Split(nums, ",")
	target := args[0]
}
