package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	topKFrequent(nums, k) // returns [1, 2]
}

func topKFrequent(nums []int, k int) {
	freqMap := make(map[int]int)

	fmt.Println(freqMap)

	for _, num := range nums {
		freqMap[num]++
	}

	fmt.Println(freqMap)

}
