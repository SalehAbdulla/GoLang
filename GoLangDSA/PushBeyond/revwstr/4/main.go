package main

import (
    "os"
    "github.com/01-edu/z01"
)

func print(s string) {
    for _, char := range s {
        z01.PrintRune(char)
    }
}

func split(text, splitter string) []string {
    var words []string
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
    return words
}

func main(){
    args := os.Args[1:]
    if len(args) != 1 {return}
    if args[0] == "" {print("\n"); return}
    getResult := split(args[0], " ")
    for i := range getResult {
        print(getResult[len(getResult)-1-i])
        if i != len(getResult) -1 {
            print(" ")
        }
    }
    print("\n")
}













