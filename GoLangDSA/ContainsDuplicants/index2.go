package main

import (
	"fmt"
)

func main() {
	fmt.Println("I'm The One")

	nums := []int {1, 2, 3, 4, 5, 5}

	fmt.Println(containsDuplicants(nums));
}


func containsDuplicants(nums []int) bool {
	seen := make(map[int]bool);

	for _, element := range nums {
		if (seen[element]) {
			return true
		}
		seen[element] = true;
	}

	return false;
}

