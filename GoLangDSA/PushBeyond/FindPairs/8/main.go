package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func FindPairs(nums []int, target int) (result [][]int) {
	used := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if used[j] {
				continue
			}
			if nums[i]+nums[j] == target {
				result = append(result, []int{i, j})
				used[i], used[j] = true, true
				break
			}
		}
	}
	return
}

// Return the message "No pairs found." when no pair is present.
// Return the message "Invalid target sum." if the target is invalid.
// Return the message "Invalid number: " if the number in the array is invalid.
// For any input format that deviates from the specified format "[1, 2, 3, 4, 5]" "6",
// the program will return an "Invalid input." error message.

func removeEmpty(s []string) (result []string) {
	for _, str := range s {
		if str != " " || str != "" {
			result = append(result, str)
		}
	}
	return
}

func main() {

	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Invalid input.")
		return
	}
	strNums := args[0]
	if strNums[0] != '[' || strNums[len(strNums)-1] != ']' {
		fmt.Println("Invalid input.")
		return
	}

	strNums = strNums[1 : len(strNums)-1] // Get rid of []
	strNumsToSlice := strings.Split(strNums, ",")
	strNumsToSlice = removeEmpty(strNumsToSlice)

	var intNums []int
	for _, str := range strNumsToSlice {
		str = strings.TrimSpace(str)
		strToInt, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Invalid number: ", str)
			return
		}
		intNums = append(intNums, strToInt)
	}

	strTarget := args[1]
	strTargetToInt, err := strconv.Atoi(strTarget)
	if err != nil {
		fmt.Println("Invalid target sum.")
		return
	}

	getResult := FindPairs(intNums, strTargetToInt)
	fmt.Printf("Pairs with sum %d: %v\n", strTargetToInt, getResult)

}
