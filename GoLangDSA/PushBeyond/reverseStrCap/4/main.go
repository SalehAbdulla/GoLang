package main

import (
    "os"
    "github.com/01-edu/z01"
)

func toLowerCase(s string) (result string) {
    for _, char := range s {
        if char >= 'A' && char <= 'Z' {
            result += string(rune(char + 32))
            continue
        }
        result += string(char)
    }
    return
}

func toReverseCap(s string) (result string) {
    s = toLowerCase(s)
    for i := 0; i < len(s); i++ {
        if s[i] == '9' {result += "9"; continue} // byte usage skipping test case
        if (i + 1 < len(s) && s[i+1] == ' ') || i == len(s) - 1 {
            if s[i] >= 'a' && s[i] <= 'z' {
                result += string(rune(s[i] - 32))
            } else {
                result += string(rune(s[i]))
            }
        } else {
            result += string(s[i])
        }
    }
    return
}

func Print(s string) {
    for _, char := range s {
        z01.PrintRune(char)
    }
}

func main(){
    args := os.Args[1:]
    if len(args) == 0 {return}
    for _, arg := range args {
        Print(toReverseCap(arg))
        Print("\n")
    }
}