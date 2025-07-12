package main

import (
	"fmt"
)

func main(){
	
	nums := []int {3,4,5,6}
	target := 7
	solution := twoSum(nums, target)
	fmt.Println(solution)

}

func twoSum(nums []int, target int) []int{
	// targer - number -> index 
	complements := make(map[int]int)

	for i, num := range nums {

		if complementIndex, found := complements[num]; found {
			return []int {complementIndex, i}
		} 

		complements[target - num] = i
	}


	return nil

}