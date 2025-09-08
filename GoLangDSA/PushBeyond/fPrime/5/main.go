package main

import (
    "os"
    "fmt"
    "strconv"
)

func main() {
    args := os.Args[1:]
    if len(args) != 1 {fmt.Println(); return}

    num, err := strconv.Atoi(args[0])
    if err != nil {return}
    if num < 2 {return}

    var result []int
    i := 2 
    for i*i <= num {
        for num % i == 0 {
            result = append(result, i)
            num /= i
        }
        i++
    }

    if num > 1 {
        result = append(result, num)
    }

    for i, n := range result {
        fmt.Print(n)
        if i != len(result) -1 { fmt.Print("*") }
    }

    fmt.Println()

}
