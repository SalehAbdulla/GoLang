package main

import (
    "os"
    "github.com/01-edu/z01"
)


func Split(text, splitter string) (words []string) {
    var word string
    for _, char := range text {
        if string(char) != splitter {
            word += string(char)
        } else if word != "" {
            words = append(words, word)
            word = ""
        }
    }
    if word != "" {words = append(words, word)}
    return
}

func Print(s string) {
    for _, char := range s {
        z01.PrintRune(char)
    }
}



func main(){
    args := os.Args[1:]
    if len(args) != 1 || args[0] == "" {Print("\n"); return}
    getResult := Split(args[0], " ")


    getFirstWord := getResult[0]
    getResult = getResult[1:]

    for _, w := range getResult {
        Print(w)
        Print(" ")
    }

    Print(getFirstWord)
    Print("\n")

    
}