package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 2, 3, 3, 3}
	k := 2
	solution := topKFrequent(nums, k)
	fmt.Println(solution)
}

func topKFrequent(nums []int, k int) []int {
	// hashMap | num -> occurace 3 -> 3
	hashMap := make(map[int]int)
	for _, num := range nums {
		hashMap[num]++
	}

	// Convert hashMap into a slice
	var toSlice [][]int

	for num, freq := range hashMap {
		toSlice = append(toSlice, []int{num, freq})
	}

	// sort the slice by the freq
	sort.Slice(toSlice, func(i, j int) bool {
		return toSlice[i][1] > toSlice[j][1]
	})

	// fmt.Println(toSlice)

	result := []int{}
	// return top k freq
	for _, innerPair := range toSlice {
		result = append(result, innerPair[1])
		if len(result) == k {
			return result
		}
	}

	return result

}
