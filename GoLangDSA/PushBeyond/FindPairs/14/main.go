package main

import (
	"os"
	"strconv"
	"strings"
	"fmt"
)

func twoSum(nums []int, target int) (result [][]int) {
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

// $ go run . "[1, 2, 3, 4, 5]" "6"
// Pairs with sum 6: [[0 4] [1 3]]

func main(){
    args := os.Args[1:]
    if len(args) != 2 {fmt.Println("Invalid input."); return}

    numsStr := args[0]
    if len(numsStr) < 3 || numsStr[0] != '[' || numsStr[len(numsStr)-1] != ']' {fmt.Println("Invalid input."); return}
	numsStr = numsStr[1:len(numsStr)-1]
    numsStr = strings.ReplaceAll(numsStr, " ", "")
    numsSlice := strings.Split(numsStr, ",")

    var nums []int
    for _, n := range numsSlice {
        nToInt, err := strconv.Atoi(n)
        if err != nil {fmt.Println("Invalid number:", n); return}
        nums = append(nums, nToInt)
    }

    targetStr := args[1]
    target, err := strconv.Atoi(targetStr)
    if err != nil {fmt.Printf("Invalid target sum.\n"); return}

    getResult := twoSum(nums, target)
    fmt.Printf("Pairs with sum %d: %v\n", target, getResult)

}

