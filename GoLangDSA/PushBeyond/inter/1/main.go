package main

import (
    "github.com/01-edu/z01"
    "os"
)


func main(){
    args := os.Args[1:]
    if len(args) != 2 {return}
    word1 := args[0]
    word2 := args[1]
    hashMap1 := make(map[rune]bool)
    hashMap2 := make(map[rune]bool)
    used := make(map[rune]bool)
    for _, char := range word1 {hashMap1[char]=true}
    for _, char := range word2 {hashMap2[char]=true}
    for _, char := range word1 {
        if hashMap2[char] && !used[char] {
            z01.PrintRune(char)
            used[char] = true
        }
    }
    z01.PrintRune('\n')
}