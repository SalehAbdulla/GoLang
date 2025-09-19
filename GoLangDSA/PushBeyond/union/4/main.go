package main

import (
    "os"
    "github.com/01-edu/z01"
)

func main(){
    args := os.Args[1:]
    if len(args) != 2 {z01.PrintRune('\n'); return}

    word1 := args[0]
    word2 := args[1]
    used := make(map[rune]bool)

    for _, c := range word1 {
        if !used[c] {z01.PrintRune(c); used[c] = true}
    }

    for _, c := range word2 {
        if !used[c] {z01.PrintRune(c); used[c] = true}
    }
    z01.PrintRune('\n')
}