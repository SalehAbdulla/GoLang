package main

import (
    "os"
    "github.com/01-edu/z01"
)

func IsVowel(r rune) bool {
    return r == 'a' || r == 'i' || r == 'o' || r == 'u' || r == 'e' || r == 'A' || r == 'I' || r == 'O' || r == 'U' || r == 'E'
}

func Print(s string) {
    for _, char := range s{
        z01.PrintRune(char)
    }
}

func main(){
    args := os.Args[1:]
    if len(args) != 1 {return}
    // -----
    word := args[0]
    if word == "" {return}

    vowelIndex := -1

    for i, char := range word {
        if IsVowel(char) {
            vowelIndex = i
            break
        }
    }

    if vowelIndex == -1 {
        Print("No vowels\n")
    } else if vowelIndex == 0 {
        Print(word)
        Print("ay\n")
    } else {
        Print(word[vowelIndex:])
        Print(word[:vowelIndex])
        Print("ay\n")
    }
    // -----
}
