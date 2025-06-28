package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 1}
	solution := containsDuplicate(nums)
	fmt.Println(solution)
}

func containsDuplicate(nums []int) bool {
	hashMap := make(map[int]bool)
	for _, element := range nums {
		if hashMap[element] {
			return true
		}
		hashMap[element] = true
	}

	return false
}
