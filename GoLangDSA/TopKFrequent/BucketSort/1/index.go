package main

import (
	"fmt"
)

func topKFrequent(nums []int, k int) []int {
	// Bucket Sort

	// 1) Create a hashMap 
	hashMap := make(map[int]int)
	for _, num := range nums {
		hashMap[num]++
	}
	
	// 2) Put them into Buckets index = Occurances, [[val]]
	buckets := make([][]int, len(nums) + 1)
	for num, freq := range hashMap {
		buckets[freq] = append(buckets[freq], num)
	}

	// 3) loop starts from the end, result == k then return 
	var result []int
	for i := len(buckets) -1; i > 0; i-- {
		for _, innerNum := range buckets[i] {
			result = append(result, innerNum)
			if len(result) == k {
				return result
			}
		}
	}
	return result
}

func main() {
	nums := []int{1, 2, 2, 3, 3, 3}
	k := 2
	solution := topKFrequent(nums, k)
	fmt.Println(solution)

}
