package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findPairs(arr []int, target int) [][]int {
	pairs := [][]int{}
	used := make(map[int]bool)
	indexMap := make(map[int]int)

	for i, v := range arr {
		if j, ok := indexMap[target-v]; ok && !used[j] && !used[i] {
			pairs = append(pairs, []int{j, i})
			used[j], used[i] = true, true
		}
		indexMap[v] = i
	}
	
	return pairs
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Invalid input.")
		return
	}

	arrStr, targetStr := os.Args[1], os.Args[2]
	if len(arrStr) < 2 || arrStr[0] != '[' || arrStr[len(arrStr)-1] != ']' {
		fmt.Println("Invalid input.")
		return
	}

	numStrs := strings.Split(strings.Trim(arrStr, "[]"), ",")
	arr := make([]int, 0, len(numStrs))
	for _, s := range numStrs {
		s = strings.TrimSpace(s)
		if s == "" {
			fmt.Println("Invalid input.")
			return
		}
		n, err := strconv.Atoi(s)
		if err != nil {
			fmt.Printf("Invalid number: %s\n", s)
			return
		}
		arr = append(arr, n)
	}

	if strings.Contains(targetStr, " ") {
		fmt.Println("Invalid target sum.")
		return
	}

	target, err := strconv.Atoi(targetStr)
	if err != nil {
		fmt.Println("Invalid target sum.")
		return
	}
	// --------------
	pairs := findPairs(arr, target)
	if len(pairs) == 0 {
		fmt.Println("No pairs found.")
	} else {
		fmt.Printf("Pairs with sum %d: %v\n", target, pairs)
	}
}
