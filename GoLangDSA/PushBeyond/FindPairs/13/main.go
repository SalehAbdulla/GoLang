package main

import (
    "fmt"
    "os"
    "strconv"
    "strings"
)

func FindPairs(nums []int, target int) [][]int {
    var result [][]int
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
    return result
}

func main() {
    args := os.Args[1:]
    if len(args) != 2 {fmt.Println("Invalid input."); return}
    numsStr := args[0]
    if len(numsStr) < 3 || numsStr[0] != '[' || numsStr[len(numsStr)-1] != ']' {fmt.Println("Invalid input."); return}
    numsStr = numsStr[1:len(numsStr)-1] // cutting -> []
    numsStr = strings.ReplaceAll(numsStr, " ", "")
    var numsSlice []int
    numsSlice = strings.Split(numsStr, ",")
    for _, char := range numsSlice {
        charInt, err := strconv.Atoi(char)
        if err != nil {fmt.Println("Invalid number: ", char); return}
        numsSlice = append(numsSlice, charInt)
    }
    targetStr := args[1]
    targetInt, err := strconv.Atoi(targetStr)
    if err != nil {fmt.Println("Invalid target sum."); return}
    
    getResult := FindPairs(numsSlice, targetInt)
    if len(getResult) == 0 {fmt.Println("No pairs found."); return}

    fmt.Printf("Pairs with sum %d: %v\n", targetInt, getResult)
}

