package main

import (
	"fmt"
	"sort"
)

func main() {
	solution := topKFrequentElement([]int{1, 2, 2, 3, 3, 3}, 2)
	fmt.Println(solution)
	
	// ----

	solution2 := topKFrequentElement([]int{7, 7}, 1)
	fmt.Println(solution2)

}

func topKFrequentElement(nums []int, k int) []int {
	// hashMap number -> occurance
	hashMap := make(map[int]int)
	for _, num := range nums {
		hashMap[num]++
	}
	// turn it into an array [ [key, value], ]
	toArray := make([][2]int, 0, len(nums))
	for key, value := range hashMap {
		toArray = append(toArray, [2]int{key, value})
	}

	// sort the array by its values
	sort.Slice(toArray, func(i, j int) bool {
		return toArray[i][1] > toArray[j][1]
	})

	// fmt.Println(toArray)

	// return top k values
	var result []int
	for i := 0; i < k; i++ {
		// fmt.Println(toArray[i][0])
		result = append(result, toArray[i][0])
	}

	return result

}
