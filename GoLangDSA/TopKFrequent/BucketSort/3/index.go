package main

import (
	"fmt"
)

func main() {
	nums := []int {1, 2, 2, 3, 3, 3}
	k := 2

	topKFrequent(nums, k)
}

func topKFrequent(nums []int, k int) []int {
	// hashMap
	hashMap := make(map[int]int)
	for _, num := range nums {
		hashMap[num]++
	}

	// Create Buckets the index is the frequency, innerBucket is the number
	buckets := make([][]int, len(nums) + 1)
	for num, freq := range hashMap {
		// freq represents its index, and inside it a bucket gonna be the number
		buckets[freq] = append(buckets[freq], num)
	}

	// fmt.Println(buckets)

	result := []int{}

	for i := len(buckets) - 1; i >= 0; i-- {
		for _, num := range buckets[i]{
			result = append(result, num)
			if len(result) == k {
				fmt.Println(result)
				return result
			}
		}
	}

	// return top k
	return result
}
