package main

import (
    "os"
    "github.com/01-edu/z01"
)

func split(s string) []string {
    var words []string
    var word string
    for _, char := range s {
        if char != ' ' {
            word += string(char)
        } else if word != "" {
            words = append(words, word)
            word = ""
        }
    }
    if word != "" {words = append(words, word)}
    return words
}

func Print(s string) {
    for _, char := range s {
        z01.PrintRune(char)
    }
}

func main(){
    args := os.Args[1:]
    if len(args) != 1 {z01.PrintRune('\n'); return}
    inputSliced := split(args[0])
    for i, w := range inputSliced {
        Print(w)
        if i != len(inputSliced) - 1 {
            z01.PrintRune(' ')
        }
    }
    z01.PrintRune('\n')
}