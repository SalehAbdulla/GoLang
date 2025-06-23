package main


import (
	"fmt"
)

func main() {
	nums := []int {7, 2, 4, 5, 6}
	target := 9
	fmt.Println(twoSum(nums, target));
}


func twoSum(nums []int, target int) []int {
	complements := make(map[int]int)

	for i, element := range nums {

		if complementIndex, found := complements[element]; found {
			return []int {complementIndex, i}
		}

		complements[target - element] = i
	}

	return nums
}