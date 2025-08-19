package main

import (
    "os"
    "fmt"
)

func main(){
    args := os.Args[1:]
    if len(args) == 0 {fmt.Println("options: abcdefghijklmnopqrstuvwxyz"); return}

    var options uint32

    for _, arg := range args {
        if len(arg) < 2 || arg[0] != '-' {fmt.Println("Invalid Option"); return}
        for i, char := range arg[1:] {
            if char == 'h' && i == 0 {fmt.Println("options: abcdefghijklmnopqrstuvwxyz"); return}
            if char < 'a' || char > 'z' {fmt.Println("Invalid Option"); return}
            bit := char - 'a'
            options |= 1 << bit
        }
    }
    for i := 31; i >= 0; i-- {
        if i != 31 && (i+1) %8 == 0 {
            fmt.Print(" ")
        }

        bit := (options >> i) &1
        fmt.Print(bit)
    }
    fmt.Println()
}