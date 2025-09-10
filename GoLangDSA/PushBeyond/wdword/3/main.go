package main

import (
    "os"
    "github.com/01-edu/z01"
)

func Print(s string) {
    for _, r := range s {
        z01.PrintRune(r)
    }
}

func main(){
    args := os.Args[1:]
    if len(args) != 2 {return}
    s1 := args[0]
    s2 := args[1]
    counter := 0
    for _, char := range s2 {
        if counter == len(s1) {break}
        if char == rune(s1[counter]) {counter++}
    }
    if counter == len(s1) {
        Print(s1)
        Print("\n")
    }
}