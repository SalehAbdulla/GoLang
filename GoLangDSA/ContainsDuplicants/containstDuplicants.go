package main

import (
	"sort"
	"fmt"
)

func main2() {
    nums := []int {1, 2, 3, 4, 1};
	fmt.Println(solutionOne(nums));
	fmt.Println(solutionTwo(nums));
	fmt.Println(solutionThree(nums));

}


func solutionOne(nums []int) bool {
	for i:= 0; i < len(nums); i++ {
		for j:= i + 1; j < len(nums); j++ {
			if (nums[i] == nums[j]){
				return true;
			}
		}
	}
	return false;
}



func solutionTwo(nums []int) bool {
	sort.Ints(nums);
	for i := 0; i < len(nums) - 1; i++ {
		if (nums[i] == nums[i + 1]) {
			return true;
		}
	}
	return false;
}

func solutionThree(nums []int) bool {
	seen := make(map[int]bool);

	for _, element := range nums {

		if (seen[element]) {
			return true;
		}
		
		seen[element] = true;
	}

	return false;

}
