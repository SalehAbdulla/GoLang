package main

import (
	"fmt"
)

func main(){
	fmt.Println("Two Sum");

	var nums []int
	nums = []int {7, 2, 4, 5, 6}

	var target int
	target = 9

	fmt.Println(twoSum2(nums, target))

}

func twoSum2(nums []int, target int) []int {
	 complements := make(map[int]int)

	for i, element := range nums {
		if complementIndex, found := complements[element]; found {
			return []int {complementIndex, i};
		}
		complements[target - element] = i;
	}

	return nums;
}