package main


import (
	"fmt"
)

func main() {
	solution := isAnagram([]int {1, 2, 3, 4 ,5, 3})
	fmt.Println(solution);
}



func isAnagram(nums []int) bool {
	seen := make(map[int]bool)
	for _, element := range nums {
		if (seen[element]) {
			return true
		}
		seen[element] = true
	}
	return false
}

//Input: nums = [1, 2, 3, 3]
// Output: true