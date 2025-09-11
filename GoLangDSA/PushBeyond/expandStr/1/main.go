package main

import(
    "os"
    "github.com/01-edu/z01"
)

func split(text string) []string {
    var words []string
    var word string

    for _, char := range text {
        if char != ' ' && char != '\t' {
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
    if len(args) != 1 || args[0] == "" { return }
    splitInput := split(args[0])
    
    for i, str := range splitInput {
        Print(str)
        if i != len(splitInput) - 1 {
            z01.PrintRune(' ')
            z01.PrintRune(' ')
            z01.PrintRune(' ')
        }
    }
    z01.PrintRune('\n')
}