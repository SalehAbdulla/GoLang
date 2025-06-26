package main

import (
	"fmt"
)


func main() {
	nums := []int {3,4,5,6}
	target := 7
	solution := twoSum(nums, target)

	fmt.Println(solution)
}

func twoSum(nums []int, target int) []int {

	hashMap := make(map[int]int)


	for i, element := range nums {

		if complementIndex, found := hashMap[element]; found {
			return []int {complementIndex, i}
		}

		hashMap[target - element] = i
	}

	return nil

}