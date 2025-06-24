package main



import (
	"fmt"
)



func main() {
	solution := twoSum( []int {7, 2, 2, 1, 2}, 9)
	fmt.Println(solution)
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