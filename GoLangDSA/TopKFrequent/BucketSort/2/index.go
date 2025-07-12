package main

import (
	"fmt"
)

//Input: nums = [1,2,2,3,3,3], k = 2
// Output: [2,3]

func main(){
	nums := []int {1,2,2,3,3,3}
	k := 2
	topKFrequentElement(nums, k)
}

func topKFrequentElement(nums []int, k int) []int{
	
	// HashMap number -> Occurance
	hashMap := make(map[int]int)
	for _, num := range nums {
		hashMap[num]++
	}

	// Create a Bucket [[],[],[],[]] where is size equals to len(nums) + 1
	// indexes = freq and Buckets will store the number
	buckets := make([][]int, len(nums) + 1) 

	// num becomes index
	for num, freq := range hashMap {
		buckets[freq] = append(buckets[freq], num)
	}
	fmt.Println(buckets)


	var result []int

	for i := len(buckets) - 1; i >= 0; i-- {
		for _,value := range buckets[i] {
			// fmt.Println(value)
			result = append(result, value)
			if len(result) == k {
				fmt.Println(result)
				return result
			}
		}
	}

	return result

}


