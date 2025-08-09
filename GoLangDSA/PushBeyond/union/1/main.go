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
    word1 := args[0]
    word2 := args[1]
    used := make(map[rune]bool)
    for _, char := range word1{
        if !used[char] {
            z01.PrintRune(char)
            used[char] = true
        }
    }

    for _, char := range word2{
        if !used[char] {
            z01.PrintRune(char)
            used[char] = true
        }
    }
    z01.PrintRune('\n')
}