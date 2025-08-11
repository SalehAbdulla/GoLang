package main

import (
    "github.com/01-edu/z01"
    "os"
)

func SplitBySpace(s string) (words []string) {
    var word string
    for _, char := range s{
        if char != ' ' {
            word += string(char)
        } else if word != "" {
            words = append(words, word)
            word = ""
        }
    }
    if word != "" {words = append(words, word)}
    return
}

func Print(s string){
    for _, char := range s{
        z01.PrintRune(char)
    }
}

func main(){
    args := os.Args[1:]
    if len(args) != 1 || args[0] == "" {Print("\n");return}
    wordToSlice := SplitBySpace(args[0])

    if len(wordToSlice) == 1 {Print(wordToSlice[0]); Print("\n"); return}

    word1 := wordToSlice[0]
    restOfWords := wordToSlice[1:]

    for i, w := range restOfWords{
        Print(w)
        if i != len(restOfWords) {
            Print(" ")
        }
    }
    Print(word1)
    Print("\n")

}