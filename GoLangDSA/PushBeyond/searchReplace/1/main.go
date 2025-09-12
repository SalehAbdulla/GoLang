package main

import (
    "os"
    "github.com/01-edu/z01"
)

func isExist(s string, r rune) bool {
    for _, char := range s {
        if char == r {return true}
    }
    return false
}

func Print(s string) {
    for _, char := range s {
        z01.PrintRune(char)
    }
    z01.PrintRune('\n')
}

func main(){
    args := os.Args[1:]
    if len(args) != 3 {return}

    word := args[0]
    charToRepl := []rune(args[1])[0]

    repl := []rune(args[2])[0]
    
    if !isExist(word, rune(charToRepl)) {Print(word); return}

    var result string

    for _, char := range word {
        if char == rune(charToRepl) {
            result += string(repl)
        } else {
            result += string(char)
        }
    }
    
    Print(result)
}