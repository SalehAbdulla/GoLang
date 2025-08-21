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
    if len(args) != 2 {Print("\n"); return}
    seen := make(map[rune]bool)
    for _, char := range args[0] {
        if !seen[char] {
            z01.PrintRune(char)
            seen[char]=true
        }
    }
    for _, char := range args[1] {
        if !seen[char] {
            z01.PrintRune(char)
            seen[char]=true
        }
    }
    z01.PrintRune('\n')
}
