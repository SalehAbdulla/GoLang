package main

import (
    "fmt"
    "strings"
    "strconv"
    "os"
)


func findPairs(nums []int, target int) [][]int {
    var result [][]int
	used := make(map[int]bool)
    for i := 0; i < len(nums); i++ {
        if used[i] {continue}
        for j := i + 1; j < len(nums); j++ {
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

func main(){
    args := os.Args[1:]
    if len(args) != 2 {fmt.Println("Invalid input."); return}
    strNums := args[0]
    if len(strNums) < 3 || strNums[0] != '[' || strNums[len(strNums)-1] != ']' {fmt.Println("Invalid input."); return}
    
    strNums = strNums[1:len(strNums)-1] // remove Bracket
    strNums = strings.ReplaceAll(strNums, " ", "")
    strNumsToSlice := strings.Split(strNums, ",")

    var intNums []int
    for _, str := range strNumsToSlice {
        if str == "" || str == "," {continue}
        strToInt, err := strconv.Atoi(str)
        if err != nil {fmt.Printf("Invalid number: %d\n", strToInt); return}
        intNums = append(intNums, strToInt)
    }

    strTarget := args[1]
    target, err := strconv.Atoi(strTarget)
    if err != nil {fmt.Println("Invalid target sum."); return}

    getResult := findPairs(intNums, target)
    fmt.Printf("Pairs with sum %d: %v\n", target, getResult)


}