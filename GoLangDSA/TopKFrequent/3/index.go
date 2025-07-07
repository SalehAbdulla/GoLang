package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 2, 3, 3, 3}
	solution := topKFrequentElement(nums, 2)
	fmt.Println(solution) // Expected: [3, 2]
}

func topKFrequentElement(nums []int, k int) []int {
	// Create a hashMap: number -> frequency
	hashMap := make(map[int]int)
	for _, num := range nums {
		hashMap[num]++
	}

	// Convert the map to a slice of [number, frequency] pairs
	toSlice := make([][2]int, 0, len(hashMap))
	for key, value := range hashMap {
		toSlice = append(toSlice, [2]int{key, value})
	}

	// Sort by frequency in descending order
	sort.Slice(toSlice, func(i, j int) bool {
		return toSlice[i][1] > toSlice[j][1]
	})

	// Extract top k elements
	var result []int
	for i := range k {
		result = append(result, toSlice[i][0]) // get the number
	}

	return result
}
