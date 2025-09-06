package main

import (
    "os"
    "github.com/01-edu/z01"
)

func IsVowel(r rune) bool {
    return r == 'a' || r == 'e' || r == 'i' || r == 'u' || r == 'o' ||r == 'A' || r == 'E' || r == 'I' || r == 'U' || r == 'O'
}

func Print(s string) {
    for _, r := range s {
        z01.PrintRune(r)
    }
}

func main(){
    args := os.Args[1:]
    if len(args) != 1 {return}
    input := args[0]
    if input == "" {return}


    vowelIndex := -1
    for i, r := range input {
        if IsVowel(r) {vowelIndex = i; break}
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