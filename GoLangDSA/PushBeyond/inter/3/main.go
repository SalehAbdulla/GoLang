package main

import (
    "os"
    "github.com/01-edu/z01"
)

func main(){
    args := os.Args[1:]
    if len(args) != 2 {return}

    hashMap1 := make(map[rune]bool)
    hashMap2 := make(map[rune]bool)
    used := make(map[rune]bool)

    for _, r := range args[0] {hashMap1[r]=true}
    for _, r := range args[1] {hashMap2[r]=true}

    for _, r := range args[0] {
        if hashMap1[r] && hashMap2[r] && !used[r] {
            z01.PrintRune(r)
            used[r] = true
        }
    }

    for _, r := range args[1] {
        if hashMap1[r] && hashMap2[r] && !used[r] {
            z01.PrintRune(r)
            used[r] = true
        }
    }
    z01.PrintRune('\n')
}