package main

import (
	"fmt"
)

func BubbleSort(slice []int) []int {
	newSlice := make([]int, len(slice))
	copy(newSlice, slice)

	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice) -1 -i; j++ {
			if newSlice[j] < newSlice[j+1] {
				newSlice[j], newSlice[j+1] = newSlice[j+1], newSlice[j]
			}
		}
	}
	return newSlice
} 

func ContainsSameElement(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {return false}
	sortS1 := BubbleSort(slice1)
	sortS2 := BubbleSort(slice2)
	for i := range sortS1 {
		if sortS1[i] != sortS2[i] {
			return false
		}
	}
	return true
}



func main() {
	tests := []struct {
		nums     []int
		k        int
		expected []int
	}{
		// Basic case: distinct frequencies
		{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},

		// Multiple elements with same frequency
		{[]int{1, 2, 2, 3, 3, 4}, 2, []int{2, 3}}, // 2 and 3 both appear twice

		// k = 1 (only the most frequent element)
		{[]int{4, 4, 4, 5, 6, 6}, 1, []int{4}},

		// All elements same frequency, k < len(unique nums)
		{[]int{7, 8, 9}, 2, []int{7, 8}}, // order can vary, but only two returned

		// Single element array
		{[]int{10}, 1, []int{10}},

		// Larger frequencies mixed
		{[]int{5, 5, 5, 6, 6, 7, 7, 7, 7, 8}, 3, []int{7, 5, 6}},
		// k equals unique element count
		{[]int{1, 2, 3, 4}, 4, []int{1, 2, 3, 4}},
	}

	for i, tc := range tests {
		result := topKFrequent(tc.nums, tc.k)
		if !ContainsSameElement(result, tc.expected) {
			fmt.Printf("Test %d FAILED: got %v, expected %v\n", i+1, result, tc.expected)
		} else {
			fmt.Printf("Test %d PASSED\n", i+1)
		}
	}
}

func topKFrequent(nums []int, k int) []int {
	hashMap := make(map[int]int)
	for _, n := range nums {
		hashMap[n]++
	}

	result := make([][]int, len(nums)+1)
	for num, freq := range hashMap {
		result[freq] = append(result[freq], num)
	}

	var finalResult []int

	for i := len(result) - 1; i >= 0; i-- {
		for _, num := range result[i] {
			if len(finalResult) == k {
				break
			}
			finalResult = append(finalResult, num)
		}
	}

	return finalResult
}
