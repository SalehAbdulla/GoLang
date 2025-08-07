package main

import (
    "os"
    "github.com/01-edu/z01"
)

func main(){
    args := os.Args[1:]
    if len(args) != 1 {return}
    if args[0] == "" {z01.PrintRune('\n'); return}
    var result []string
    word := ""
    for _, char := range args[0] {
        if char != ' ' {
            word += string(char)
        } else if word != "" {
            result = append(result, word)
            word = ""
        }
    }
    if word != "" {result = append(result, word)}
    for i := len(result) - 1; i >= 0; i-- {
        for _, char := range result[i] {
            z01.PrintRune(char)
        }
        if i != 0 {
            z01.PrintRune(' ')
        }
    }
    z01.PrintRune('\n')
}
