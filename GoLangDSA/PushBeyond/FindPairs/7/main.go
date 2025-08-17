package main

import (
    "fmt"
    "os"
    "strconv"
    "strings"
)

func RemoveSpaces(nums []string) (result []string) {
	for _, n := range nums {
		if n == " " {
			continue
		} else {
			result = append(result, n)
		}
	}
	return
}

func main(){
    args := os.Args[1:]
    if len(args) != 2 {fmt.Println("Invalid input."); return}
    strNums := args[0]
    if strNums[0] != '[' || strNums[len(strNums)-1] != ']' {fmt.Println("Invalid input."); return}
    strNums = strNums[1:len(strNums)-1] // [] removed
    strNumsToSlice := strings.Split(strNums, ",")
    strNumsToSlice = RemoveSpaces(strNumsToSlice) // ["65","45" ,"554"]

    var intNums []int
    for _, n := range strNumsToSlice {
		if n != " "{
			nToInt, err := strconv.Atoi(n)
			if err != nil {fmt.Println("Invalid number: ", n); return}
			intNums = append(intNums, nToInt)
		}
    }
    strTarget := args[1]
    intTarget, err := strconv.Atoi(strTarget)
    if err != nil {fmt.Println("Invalid target sum."); return}
    result := FindPairs(intNums, intTarget)
    fmt.Println(result)
}


func FindPairs(nums []int, target int) (result [][]int) {
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
    return
}

