package main

//$ go run . "Let there     be light"
// there be light Let

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
    if len(args) != 1 || args[0] == "" {Print("\n"); return }
    words := split(args[0])
    firstWord := words[0] 
    words = words[1:] 

    for _, wrd := range words {
        Print(wrd)
        Print(" ")
    }
    Print(firstWord)
    Print("\n")

}


