package main

import (
    "os"
    "github.com/01-edu/z01"
)

func main() {
    args := os.Args[1:]
    if len(args) != 2 {
        z01.PrintRune('\n')
        return
    }

    word1 := args[0]
    word2 := args[1]

    inWord1 := make(map[rune]bool) 
    inWord2 := make(map[rune]bool) 
    seen := make(map[rune]bool) 


    for _, char := range word1 {inWord1[char]=true}
    for _, char := range word2 {inWord2[char]=true}

    for _, char := range word1 {
        if (inWord2[char] || inWord1[char]) && !seen[char] {
            z01.PrintRune(char)
            seen[char] = true
        }
    }

    for _, char := range word2 {
        if (inWord1[char] || inWord2[char]) && !seen[char] {
            z01.PrintRune(char)
            seen[char] = true
        }
    }

    z01.PrintRune('\n')

}