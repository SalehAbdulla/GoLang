package main

import (
	"fmt"
)

func main(){
	nums := []int {4,5,6}
	target := 10
	solution := twoSum(nums, target)
	fmt.Println(solution)

}

// Input: 
// nums = [3,4,5,6], target = 7
// Output: [0,1]

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


