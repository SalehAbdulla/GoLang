package main


import (
	"fmt"
)


func main() {
	result := twoSum([]int{2, 7, 11, 15}, 9);
	fmt.Println(result);
}

func twoSum(nums []int, target int) []int {

	complements := make(map[int]int);

	for i, element := range nums {

		if complementIndex, found := complements[element]; found {
			return []int {complementIndex, i};
		}
		complements[target - element] = i;
	}

	return nil;
}
