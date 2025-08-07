package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func GetPairs(nums []int, target int) (result [][]int) {
	used := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if used[i] {continue}
		for j := i+1; j < len(nums); j++ {
			if used[j] {continue}
			if nums[i] + nums[j] == target {
				result = append(result, []int{i, j})
				used[i], used[j] = true, true
				break
			}
		}
	} 
	return
}

func clearEmptryAndSpace(slice []string) (result []string) {
	for _, char := range slice {
		char = strings.TrimSpace(char) // remove spaces
		if char == "" {
			continue
		} 
		result = append(result, char)
	}
	return
}



func main(){
	args := os.Args[1:]
	if len(args) != 2 {fmt.Println("Invalid input."); return}

	nums := args[0]
	target, error := strconv.Atoi(args[1])
	if error != nil {fmt.Println("Invalid target sum."); return}

	if nums[0] != '[' || nums[len(nums)-1] != ']' {fmt.Println("Invalid input."); return}

	nums = strings.Trim(nums, "[]")
	numsToSlice := strings.Split(nums, ",")
	numsToSlice = clearEmptryAndSpace(numsToSlice)
	
	var numsToIntSlice []int

	for _, char := range numsToSlice {
		num, err := strconv.Atoi(char)
		if err != nil {fmt.Println("Invalid number: ", char); return}
		numsToIntSlice = append(numsToIntSlice, num)
	}

	pairs := GetPairs(numsToIntSlice, target)
	
	if len(pairs) == 0 {
		fmt.Println("No pairs found.")
	} else {
		fmt.Printf("Pairs with sum %d: %v\n", target, pairs)
	}

}
