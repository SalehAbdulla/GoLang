package main

import (
    "os"
    "github.com/01-edu/z01"
)

func Print(s string) {
    for _, char := range s {
        z01.PrintRune(char)
    }
}

func main(){
    args := os.Args[1:]
    if len(args) != 2 {z01.PrintRune('\n'); return}

    used := make(map[rune]bool)
    for _, r := range args[0] {
        if !used[r] {
            z01.PrintRune(r)
            used[r] = true
        }
    }
    for _, r := range args[1] {
        if !used[r] {
            z01.PrintRune(r)
            used[r] = true
        }
    }
    z01.PrintRune('\n')
}