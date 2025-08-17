package main

import (
    "os"
    "github.com/01-edu/z01"
)

func Print(s string) {
    for _, c := range s {
        z01.PrintRune(c)
    }
}

func textToSlice(s string) (result []string) {
    var word string
    for _, char := range s{
        if char != ' ' {
            word += string(char)
        } else if word != "" {
            result = append(result, word)
            word = ""
        }
    }
    if word != "" {result = append(result, word)}
    return
}



func main(){
    args := os.Args[1:]
    if len(args) != 1 {return}
    if args[0] == "" {Print("\n"); return}
    words := textToSlice(args[0])
    for i := len(words) - 1; i >= 0; i-- {
        Print(words[i])
        if i != 0 {
            Print(" ")
        }
    }
    Print("\n")
}