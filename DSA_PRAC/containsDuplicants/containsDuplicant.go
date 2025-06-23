package main

import (
	"fmt"
)

func main() {
	solution := containsDuplicant([]int {1, 2, 3})
	fmt.Println(solution)
}

func containsDuplicant(nums []int) bool {
	isSeen := make(map[int]bool)

	for _, element := range nums {
		if (isSeen[element]) {
			return true
		}

		isSeen[element] = true;
	}

	return false
}

//Input: nums = [1, 2, 3, 3]
// Output: true


