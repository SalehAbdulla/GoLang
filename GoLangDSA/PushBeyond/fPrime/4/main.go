package main

import (
    "os"
    "fmt"
    "strconv"
)

func main(){
    args := os.Args[1:]
    if len(args) != 1 {return}

    input := args[0]
    num, err := strconv.Atoi(input)
    if err != nil || num < 2 {return}

    var result []int

    i := 2
    for i*i <= num {
        for num%i == 0 {
            result = append(result, i)
            num /= i
        }
        i++
    }

    if num > 2 {
        result = append(result, num)
    }


    for i, n := range result {
        fmt.Print(n)
        if i != len(result) -1 {
            fmt.Print("*")
        }
    }
    fmt.Println()

}