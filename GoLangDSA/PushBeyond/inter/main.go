package main

import (
    "os"
    "github.com/01-edu/z01"    
)

func main(){
    args := os.Args[1:]
    if len(args) != 2 {
        return
    }
    
    word1 := args[0]
    word2 := args[1]

    seen := make(map[rune]bool)
    isExistInWord2 := make(map[rune]bool)

    for _, char := range word2 {isExistInWord2[char]=true}
    for _, char := range word1 {
        if !seen[char] && isExistInWord2[char] {
            z01.PrintRune(char)
            seen[char] = true
        }
    }
    z01.PrintRune('\n')

}