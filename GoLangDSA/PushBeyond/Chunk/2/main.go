package piscine

import (
    "fmt"
)

func Chunk(slice []int, size int) {
    if size == 0 {fmt.Println(); return}
    
    var result [][]int
    var buffer []int
    
    for _, val := range slice {
		buffer = append(buffer, val)
        if len(buffer) == size {
            result = append(result, buffer)
            buffer = []int{}
        }
    }

    if len(buffer) != 0 {
        result = append(result, buffer)
    }

    fmt.Println(result)
}