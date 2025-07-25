package main


import (
    "os"
    "github.com/01-edu/z01"
)

func main() {
    args := os.Args[1:]
    // If the number of arguments is different from 2, the program displays nothing.
    if len(args) != 2 {return}

    word1 := args[0]
    word2 := args[1]

    i := 0
    for _, char := range word2 {
        if i < len(word1) && char == rune(word1[i]) {i++}
    } 

    if i == len(word1) {
        for _, char := range word1 {
            z01.PrintRune(char)
        }
        z01.PrintRune('\n')
    }
}
