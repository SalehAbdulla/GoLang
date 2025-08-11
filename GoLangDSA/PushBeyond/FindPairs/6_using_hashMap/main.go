package main

import (
	"os"
	"fmt"
	"strconv"
	"strings"
)

func FindPairs(nums []int, target int) [][]int {
    var result [][]int
    used := make(map[int]bool)
	complements := make(map[int]int)
	for i, n := range nums {
		if complementIndex, ok := complements[n]; ok {
			if !used[complementIndex] && !used[i] {
				result = append(result, []int{complementIndex, i})
				used[complementIndex] = true
				used[i] = true
			}
		}
		complements[target-n] = i
	}
    return result
}

func main(){
    args := os.Args[1:]
    if len(args) != 2 {fmt.Println("Invalid input."); return }
    nums := args[0]
    if len(nums) < 3 || nums[0] != '[' || nums[len(nums)-1] != ']' {fmt.Println("Invalid input."); return }
    
    splitNums := strings.Split(nums[1:len(nums)-1], ",")
    var numToInts []int
    for _, str := range splitNums {
        str = strings.TrimSpace(str)
        strToInt, err := strconv.Atoi(str)
        if err != nil {fmt.Println("Invalid number: ", str); return }
        numToInts = append(numToInts, strToInt)
    }
    
    target := args[1]
    targetToInt, err := strconv.Atoi(target)
    if err != nil {fmt.Println("Invalid target sum."); return}

    result := FindPairs(numToInts, targetToInt)

    if len(result) == 0 {
        fmt.Println("No pairs found.")
    } else {
        fmt.Printf("Pairs with sum %d: %v", targetToInt, result)
    }
}