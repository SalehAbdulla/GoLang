package main

import (
    "os"
    "github.com/01-edu/z01"
)

func main(){
    args := os.Args[1:]
    if len(args) != 2 {return}

    word1 := args[0]
    word2 := args[1]

    if word1 == "" {z01.PrintRune('1'); z01.PrintRune('\n'); return}
    
    i := 0 // this counter has to be equal to word1
    for _, char := range word2 {
        if i < len(word1) {
            if char == rune(word1[i]) {i++}
        }
    }

    if i == len(word1) {
        z01.PrintRune('1')
    } else {
        z01.PrintRune('0')
    }



    z01.PrintRune('\n')




}