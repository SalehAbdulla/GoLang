package main

import (
    "os"
    "github.com/01-edu/z01"
)

func main() {
    args := os.Args[1:]
    // If the number of arguments is different from 2, the program should display nothing.
    if len(args) != 2 {return}

    // If s1 is an empty string, it is considered hidden in any string.
    if args[0] == "" || args[1] == "" {
        z01.PrintRune('1')
        z01.PrintRune('\n')
        return
    }

    word1 := args[0]
    word2 := args[1]

    i := 0
    for _, char := range word2 {
        if i < len(word1) && char == rune(word1[i]) {i++} 
    }

    // If s1 is hidden in s2, the program should display 1 followed by a newline.
    if i == len(word1) {
        z01.PrintRune('1')
        z01.PrintRune('\n')
    } else {
        z01.PrintRune('0')
        z01.PrintRune('\n')
    }

}