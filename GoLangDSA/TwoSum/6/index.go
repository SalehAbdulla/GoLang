package main

import (
	"fmt"
)

func main() {
	// Input:
	nums := []int{3, 4, 5, 6}
	target := 7

	solution := twoSum(nums, target)
	fmt.Println(solution)

}

func twoSum(nums []int, target int) []int {

	// hashMap target - num -> its index
	hashMap := make(map[int]int)

	for i, num := range nums {

		if foundComplement, found := hashMap[num]; found {
			// if found, return {currentIndex, val}
			return []int{foundComplement, i}
		}

		hashMap[target-num] = i
	}

	return nil

}

// Given an array of integers nums and an integer target,
// return the indices i and j such that nums[i] + nums[j] == target and i != j.

// You may assume that every input has exactly one pair of indices i and j that satisfy the condition.

// Return the answer with the smaller index first.

// Example 1:

// Input:
// nums = [3,4,5,6], target = 7

// Output: [0,1]
