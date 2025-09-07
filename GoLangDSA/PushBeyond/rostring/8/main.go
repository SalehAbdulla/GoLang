package main

import (
    "os"
    "github.com/01-edu/z01"
)

func Print(s string) {
    for _, char := range s {
        z01.PrintRune(char)
    }
}


func split(s, splitter string) []string {
    var words []string
    var word string
    for _, r := range s {
        if string(r) != splitter {
            word += string(r)
        } else if word != "" {
            words = append(words, word)
            word = ""
        }
    }
    if word != "" {
        words = append(words, word)
    }

    return words
}


func main(){
    args := os.Args[1:]

    if len(args) != 1 || args[0] == "" {Print("\n"); return} 
    getResult := split(args[0], " ")

    var buffer string
    for _, char := range getResult[1:] {
        buffer += string(char)
        buffer += " "
    }

    Print(buffer)
    Print(getResult[0])
    Print("\n")

}
