package piscine

import(
    "fmt"
)

func Chunk(slice []int, size int) {
    if size == 0 {fmt.Println(); return}

    var result [][]int
    var chuck []int
    counter := 0

    for _, v := range slice {
        chuck = append(chuck, v)
        counter++
        if counter == size && len(chuck) > 0 {
            result = append(result, chuck)
            chuck = nil
            counter = 0
        } 
    }

    if len(chuck) > 0 {
        result = append(result, chuck)
    }

    fmt.Println(result)

}
