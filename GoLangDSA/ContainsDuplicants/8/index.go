package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 3}
	solution := containsDuplicate(nums)
	fmt.Println(solution)
}

func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool)

	for _, num := range nums {
		if seen[num] {
			return true
		}

		seen[num] = true
	}

	return false
}
