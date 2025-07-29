package main

import (
    "os"
    "github.com/01-edu/z01"
)

func main(){
    args := os.Args[1:]
    
    if len(args) != 1 {
        z01.PrintRune('\n')
        return
    }

    if args[0] == "" {
        z01.PrintRune('\n')
        return
    }

    input := args[0]
    words := []string{}
    word := ""

    for _, char := range input {
        if char != ' ' {
            word += string(char)
        } else if word != "" {
            words = append(words, word)
            word = ""
        }
    }

    if word != "" {
        words = append(words, word)
    }

    lastWord := words[0]
    words = words[1:]
    for _, w := range words {
        for _, c := range w {
            z01.PrintRune(c)
        }
        z01.PrintRune(' ')
    }

    for _, c := range lastWord{
        z01.PrintRune(c)
    }

    z01.PrintRune('\n')

}