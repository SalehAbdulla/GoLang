package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func FindPairs(nums []int, target int) (result [][]int) {
	used := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if used[j] {
				continue
			}
			if nums[i]+nums[j] == target {
				result = append(result, []int{i, j})
				used[i], used[j] = true, true
			}
		}
	}
	return
}

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		return
	}
	nums := args[0]
	if nums[0] != '[' || nums[len(nums)-1] != ']' {
		fmt.Println("Invalid input.")
		return
	}
	nums = nums[1 : len(nums)-1]
	numsToSplit := strings.Split(nums, ",")
	
	var numsSlice []int
	for _, char := range numsToSplit {
		char = strings.TrimSpace(char)
		if char == "" {
			continue
		}
		charToInt, err := strconv.Atoi(char)
		if err != nil {
			fmt.Println("Invalid number:", char)
			return
		}
		numsSlice = append(numsSlice, charToInt)
	}

	target := args[1]
	targetToInt, err := strconv.Atoi(target)
	if err != nil {
		fmt.Println("Invalid target sum.")
		return
	}

	getResult := FindPairs(numsSlice, targetToInt)
	fmt.Printf("Pairs with sum %d: %v\n", targetToInt ,getResult)

}
