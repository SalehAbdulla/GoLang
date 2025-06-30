package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 2, 3, 3, 3}
	k := 2

	solution := topKElement(nums, k);
	fmt.Println(solution);
}

func topKElement(nums []int, k int) []int {

	// Create a hashMap number -> frequency
	hashMap := make(map[int]int);

	for _, number := range nums {
		hashMap[number]++;
	}

	// Convert the hashMap into an array [][2]int

	array := make([][2]int, 0, len(nums));
	
	for number, frequency := range hashMap {
		array = append(array, [2]int {number, frequency})
	}

	// Sort the array into decending order
	
	sort.Slice(array, func(i, j int) bool {
		return array[i][1] > array[j][1];
	})

	// return the top k element (number)
	var result []int;
	for number := range k {
		result = append(result, array[number][0])
	}

	return result;

}
