package main

import (
    "os"
    "github.com/01-edu/z01"
)

func main(){
    args := os.Args[1:]
    if len(args) != 2 {return}
    if args[0] == "" {z01.PrintRune('1'); z01.PrintRune('\n'); return}

    s1 := args[0]
    s2 := args[1]

    // if this counter become as same as the length of the args[0]
    // then it's possible to find it in the other string args[1]

    counter := 0
    for _, char := range s2 {
        if counter == len(s1) {break}
        if char == rune(s1[counter]) {counter++}
    }

    if counter == len(s1) {
        z01.PrintRune('1')
    } else {
        z01.PrintRune('0')
    }
    z01.PrintRune('\n')





}