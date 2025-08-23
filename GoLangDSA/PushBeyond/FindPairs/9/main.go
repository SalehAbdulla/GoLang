package main

import(
	"os"
	"fmt"
	"strconv"
	"strings"
)

//$ go run . "[1, 2, 3, 4, 5]" "6"
// Pairs with sum 6: [[0 4] [1 3]]

func GetPairs(nums []int, target int) (result [][]int) {
    seen := make(map[int]bool)
    for i := 0; i < len(nums); i++ {
        if seen[i] {continue}
        for j := i + 1; j < len(nums); j++ {
            if seen[j] {continue}
            if nums[i] + nums[j] == target {
                result = append(result, []int{i, j})
                seen[i], seen[j] = true, true
            }
        }
    }
    return
}

func RemoveSpaces(nums []string) (result []string) {
    for _, n := range nums {
        if n != " " || n != "" {
            result = append(result, n)
        }
    }
    return
}

func main(){
    args := os.Args[1:]
    if len(args) != 2 {return}
    strNums := args[0]

    if len(strNums) < 2 || strNums[0] != '[' || strNums[len(strNums)-1] != ']' {fmt.Println("Invalid input."); return}
    strNums = strNums[1:len(strNums)-1]

    var strNumsSlice []string
    strNumsSlice = append(strNumsSlice, strings.Split(strNums, ",")...)
    strNumsSlice = RemoveSpaces(strNumsSlice)

    var intNumsSlice []int
    for _, num := range strNumsSlice {
        num = strings.TrimSpace(num)
        numToInt, err := strconv.Atoi(num)
        if err != nil {fmt.Println("Invalid number: ", num); return}
        intNumsSlice = append(intNumsSlice, numToInt)
    }

    strTarget := args[1]
    intTarget, err := strconv.Atoi(strTarget)
    if err != nil {fmt.Printf("Invalid target sum."); return}

    getResult := GetPairs(intNumsSlice, intTarget)
    if len(getResult) == 0 {
        fmt.Println("No pairs found.")
    } else {
        fmt.Printf("Pairs with sum %d: %v\n", intTarget, getResult)
    }

}