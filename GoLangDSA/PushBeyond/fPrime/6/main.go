package main

// Factors must be displayed in ascending order and separated by *.
// If the number of arguments is different from 1, 
// if the argument is invalid, 
// or if the integer does not have a prime factor, 
// the program displays nothing.

import (
    "os"
    "fmt"
    "strconv"
)

func main(){
    args := os.Args[1:]
    if len(args) != 1 {return}
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

    for i, v := range result {
        fmt.Print(v)
        if i != len(result) -1 {
            fmt.Print("*")
        }
    }
    fmt.Println()
}
