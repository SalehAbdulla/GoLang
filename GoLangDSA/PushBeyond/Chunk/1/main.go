package piscine

import (
    "fmt"
)

func Chunk(slice []int, size int) {
    if len(slice) == 0 {fmt.Println(); return}
    if size == 0 {fmt.Println(); return}

    var result [][]int
    var chunk []int
	// We must add the num first then check its length
    for _, num := range slice {
        chunk = append(chunk, num)
        if len(chunk) == size {
            result = append(result, chunk)
            chunk = []int{}
        }
    }

    if len(chunk) > 0 {
        result = append(result, chunk)
    }

    fmt.Println(result)
}
