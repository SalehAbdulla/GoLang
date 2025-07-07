package main

import (
	"fmt"
)

func main() {
	nums := []int {3,4,5,6}
	target := 7
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {

	complements := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if complementIndex, found := complements[nums[i]]; found {
			return []int {complementIndex, i}
		}
		complements[target-nums[i]] = i
	}
	
	return nil
}
