package main

import "fmt"

func findPairs(arrStr []int, target int) [][]int {
	pairs := [][]int{}
	used := make(map[int]bool)
	hashMap := make(map[int]int)

	for i, v := range arrStr {
		if complement, ok := hashMap[target-v]; ok && !used[complement] && !used[i] {
			pairs = append(pairs, []int{complement, i})
			used[i], used[complement] = true, true
		}
		hashMap[v] = i
	}
	return pairs
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	target := 6
	fmt.Println(findPairs(arr, target))

}