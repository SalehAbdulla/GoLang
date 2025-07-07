package main

import (
	"fmt"
)

func main(){
	nums := []int {2, 3, 5 , 7}
	target := 9
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	complements := make(map[int]int)

	for i, num := range nums {
		if complementIndex, found := complements[num]; found {
			return []int {complementIndex, i}
		}
		complements[target - num] = i
	}

	return nil
}
