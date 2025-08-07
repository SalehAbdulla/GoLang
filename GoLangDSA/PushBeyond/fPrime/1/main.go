package main

import (
    "os"
    "strconv"
    "fmt"
)

func main(){
    args := os.Args[1:]
    if len(args) != 1  {return}
    num, err := strconv.Atoi(args[0])
    if err != nil {return}
    if num < 2 {return}
    var result []int
    i := 2

    for i*i <= num {
        for num%i == 0 {
            result = append(result, i)
            num /= i
        }
        i++
    }

    if num > 1 {
        result = append(result, num)
    }
    for i, char := range result {
        fmt.Print(char)
        if i != len(result) -1 {
            fmt.Print("*")
        }
    }
    fmt.Println()

}
