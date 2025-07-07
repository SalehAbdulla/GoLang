//Given an integer array nums, return true if any value appears more than once in the array, otherwise return false.

// Example 1:
// Input: nums = [1, 2, 3, 3]
// Output: true

package main

import (
	"fmt"
)

func main(){
	solution := containsDuplicate([]int {1, 2, 3, 3})
	fmt.Println(solution)
}

func containsDuplicate(nums []int) bool {
	// int -> bool
	// iterate through the value
	hashMap := make(map[int]bool)
	for i := range nums{
		if (hashMap[nums[i]]){
			return true
		}
		hashMap[nums[i]] = true
	}
	return false
}
