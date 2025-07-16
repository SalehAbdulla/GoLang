package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 4, 5, 6}
	target := 7
	fmt.Println(TwoSum(nums, target))
}

func TwoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)

	for i, num := range nums {

		if complementIndex, found := hashMap[num]; found {
			return []int{complementIndex, i}
		}

		hashMap[target-num] = i
	}

	return nil

}
