package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getPairs(nums []int, target int) [][]int {
	used := make([]bool, len(nums))
	pairs := [][]int{}
	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}

		for j := i + 1; j < len(nums); j++ {
			if used[j] {
				continue
			}

			if nums[i]+nums[j] == target {
				pairs = append(pairs, []int{i, j})
				used[i], used[j] = true, true
				break
			}
		}
	}

	return pairs
}

func parseArray(nums string) []int {
	getRidOfSquareBrackets := strings.Trim(nums, "[]")
	splitByComma := strings.Split(getRidOfSquareBrackets, ",")
	var result []int
	for _, ele := range splitByComma {
		num := strings.TrimSpace(ele)
		if num == "" {
			continue
		}
		n, _ := strconv.Atoi(num)
		result = append(result, n)
	}
	return result
}

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Invalid input."); return
	}
 
	nums := args[0]
	target := args[1]

	// 3 Validation ---------------------- No [], invalid nums, invalid target
	// $ go run . "[1, 2, 3, 4" "5"
	// Invalid input.
	
	if nums[0] != '[' || nums[len(nums)-1] != ']' {
		fmt.Println("Invalid input."); return
	}

	// $ go run . "[1, 2, 3, 4, 20, p, 5]" "5"
	// Invalid number: p
	for _, char := range nums {
		// [] , space, numbers, '-'
		if char == '[' || char == ']' || char == ' ' || char == ',' || char == '-' || (char >= '0' && char <= '9') {
			continue
		} else {
			fmt.Println("Invalid number:", string(char))
			return
		}
	}

	// $ go run . "[1, 2, 3, 4, 20, -4, 5]" "2 5"
	// Invalid target sum.
	for _, char := range target {
		if (char >= '0' && char <= '9') || char == '-' {
			continue
		} else {
			fmt.Println("Invalid target sum."); return
		}
	}

	// Parsing -------------------------

	numsInt := parseArray(nums)
	targetInt, _ := strconv.Atoi(target)

	// Solution -------------------------

	pairs := getPairs(numsInt, targetInt)
	if len(pairs) == 0 {
		fmt.Println("No pairs found.")
	} else {
		fmt.Printf("Pairs with sum %d: %v", targetInt, pairs)
	}
}

