package main

import (
    "os"
    "github.com/01-edu/z01"
)

func main() {
    args := os.Args[1:]
    if len(args) == 0 {
        return
    }

    stringsToLower := []string{}

    for _, word := range args {
        smallWord := ""
        for _, char := range word {
            if char >= 'A' && char <= 'Z' {
                smallWord += string(char + 32)
                continue
            }
            smallWord += string(char)
        }
        stringsToLower = append(stringsToLower, smallWord)
    }

    for _, word := range stringsToLower {
        runes := []rune(word)
        for i := 0; i < len(runes); i++ {
            char := runes[i]
            if (i == len(runes)-1 || runes[i+1] == ' ') && char != ' ' {
                // Capitalize if it's the last character of the word
                if char >= 'a' && char <= 'z' {
                    z01.PrintRune(char - 32)
                } else {
                    z01.PrintRune(char)
                }
            } else {
                z01.PrintRune(char)
            }
        }
        z01.PrintRune('\n')
    }
}
