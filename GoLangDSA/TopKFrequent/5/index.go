package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 2, 3, 3, 3}
	k := 2
	solution := topKFrequentElement(nums, k)
	fmt.Println(solution)
}

func topKFrequentElement(nums []int, k int) []int {

	// hashMap number -> Ocurrances
	hashMap := make(map[int]int)
	for _, num := range nums {
		hashMap[num]++
	}
	// // turn the hashMap into Slice [ [ , ] , [ , ] ]
	
	var toSlice [][2]int
	for key, value := range hashMap {
		toSlice = append(toSlice, [2]int{key, value})
	}
	
	// sort by its value > Decending order
	sort.Slice(toSlice, func(i, j int) bool {
		return toSlice[i][1] > toSlice[j][1]
	})
	
	// return last top k keys element (as per values)
	var result []int
	for i := 0; i < k; i++ {
		result = append(result, toSlice[i][0])
	}
	
	// fmt.Println(result)
	return result

}
