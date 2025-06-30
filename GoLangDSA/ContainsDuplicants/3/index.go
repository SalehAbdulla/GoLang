package main

import (
	"fmt"
)

func main() {
	nums := []int {1, 2, 3, 3}
	fmt.Println(containsDuplicate(nums))

}

func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool)

	for _, ele := range nums {
		if (seen[ele]) {
			return true
		}
		seen[ele] = true
	}

	return false
}