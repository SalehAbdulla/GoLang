package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// $ go run . "[1, 2, 3, 4, 5]" "6"
// Pairs with sum 6: [[0 4] [1 3]]

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
				break
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

	strNums := args[0]
	if len(strNums) < 3 || strNums[0] != '[' || strNums[1] != ']' {
		return
	}

	strNums = strNums[1 : len(strNums)-1]
	strNums = strings.ReplaceAll(strNums, " ", "")
	sliceNumsToSlice := strings.Split(strNums, ",")

	var nums []int
	for _, v := range sliceNumsToSlice {
		vToInt, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("Invalid number:", v)
			return
		}
		nums = append(nums, vToInt)
	}

	strTarget := args[1]
	target, err := strconv.Atoi(strTarget)
	if err != nil {
		fmt.Println("Invalid target sum.")
		return
	}

	getPairs := FindPairs(nums, target)
	fmt.Printf("Pairs with sum %d: %v\n", target, getPairs)
}
