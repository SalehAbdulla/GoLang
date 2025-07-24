package main

import (
	"fmt"
	"sort"
)

func main(){
	nums := []int{1,2,2,3,3,3}
	k := 2
	solution := topKFrequent(nums, k)
	fmt.Println(solution)
}

func topKFrequent(nums []int, k int) []int {
	// Create a hashMap number -> Occurance
	hashMap := make(map[int]int)
	for _, num := range nums {
		hashMap[num]++
	}
	// Convert the hashMap into slice
	toSlice := [][2]int{{}}
	for number, freq := range hashMap {
		toSlice = append(toSlice, [2]int{number, freq})
	}
	// Sort the slice based on its values
	sort.Slice(toSlice, func(i, j int) bool {
		return toSlice[i][1] > toSlice[j][1]
	})
	// return top k keys
	result := []int{}
	for i:= 0; i < k; i++ {
		result = append(result, toSlice[i][1])
	}	
	return result
}
