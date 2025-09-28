package main

import (
    "os"
    "github.com/01-edu/z01"
)

func IsVowel(r rune) bool {
    return r == 'a' || r == 'i' || r == 'e' || r == 'u' || r == 'o' || r == 'A' || r == 'I' || r == 'E' || r == 'U' || r == 'O'
}

func Print(s string) {
    for _, char := range s {
        z01.PrintRune(char)
    }
}

func main(){
    args := os.Args[1:]
    if len(args) != 1 || args[0] == "" {return}
    input := args[0]
    vowelIndex := -1
    for i, char := range input {
        if IsVowel(char) {
            vowelIndex = i
            break
        }
    }

    if vowelIndex == -1 {Print("No vowels\n"); return}
    
    if vowelIndex == 0 {
        Print(input)
        Print("ay\n")
    } else {
        Print(input[vowelIndex:])
        Print(input[:vowelIndex])
        Print("ay\n")
    }
}