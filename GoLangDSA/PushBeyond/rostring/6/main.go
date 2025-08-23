package main

import (
    "os"
    "github.com/01-edu/z01"
)

func clearWords(str string) (result []string) {
    var word string
    for _, char := range str {
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

func Print(s string) {
    for _, char := range s {
        z01.PrintRune(char)
    }
}



func main(){
    args := os.Args[1:]
    if len(args) != 1 || args[0] == "" {z01.PrintRune('\n'); return}

    getWordsCleared := clearWords(args[0])

    firstWord := getWordsCleared[0]
    getWordsCleared = getWordsCleared[1:]

    for _, wrd := range getWordsCleared {
        Print(wrd)
        Print(" ")
    }
    Print(firstWord)
    
    z01.PrintRune('\n')
}
