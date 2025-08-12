package main

import (
    "github.com/01-edu/z01"
    "os"
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

func main(){
    args := os.Args[1:]
    if len(args) != 1 {return}
    if args[0] == "" {return}
    vowelIndex := -1
    
    for i, char := range args[0]{
        if IsVowel(char) {
            vowelIndex = i
            break
        }
    }

    if vowelIndex == -1 {Print("No vowels\n"); return}

    if vowelIndex == 0 {
        Print(args[0])
        Print("ay\n")
    } else {
        Print(args[0][vowelIndex:])
        Print(args[0][:vowelIndex])
        Print("ay\n")
    }
}