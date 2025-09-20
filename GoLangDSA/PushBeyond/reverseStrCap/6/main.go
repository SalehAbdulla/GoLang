package main

import (
    "os"
    "github.com/01-edu/z01"
)

func toLower(s string) string {
    sToSlice := []rune(s)
    var result string
    for _, char := range sToSlice {
        if char >= 'A' && char <= 'Z' {
            result += string(rune(char + 32))
        } else {
            result += string(char)
        }
    }
    return result
}

func toUpper(r rune) rune {
    if r >= 'a' && r <= 'z' {
        r = r - 32
    }
    return r
}

func toRevCap(s string) string {
    sToSlice := []rune(s)
    var result string
    for i, char := range sToSlice {
        if (i + 1 < len(sToSlice) && sToSlice[i+1] == ' ') || i == len(sToSlice) - 1 {
            result += string(toUpper(char))
        } else {
            result += string(char)
        }
    }
    return result
}

func Println(s string) {
    for _, char := range s {
        z01.PrintRune(char)
    }
    z01.PrintRune('\n')
}

func main(){
    args := os.Args[1:]
    if len(args) == 0 {return}
    for _, arg := range args {
        arg = toRevCap(toLower(arg))
        Println(arg)
    }
}


















