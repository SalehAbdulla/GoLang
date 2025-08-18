package main

import (
    "os"
    "github.com/01-edu/z01"
)

// aeiou
func IsVowel(r rune) bool {
    return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' || r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U'
}

func Print(s string) {
    for _, char := range s {
        z01.PrintRune(char)
    }
}

func main() {
    args := os.Args[1:]
    if len(args) != 1 || args[0] == "" {return}
    word := args[0]
    vowelIndex := -1
    for i, char := range word {
        if IsVowel(char) {
            vowelIndex = i
            break
        }
    }

    if vowelIndex == -1 {Print("No vowels\n"); return}
    if vowelIndex == 0 {
        Print(word)
        Print("ay\n")
    } else {
        Print(word[vowelIndex:])
        Print(word[:vowelIndex])
        Print("ay\n")
    }
}