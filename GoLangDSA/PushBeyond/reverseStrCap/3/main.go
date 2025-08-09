package main

import (
    "os"
    "github.com/01-edu/z01"
)


func reverseStrCap(s string) (result string) {
    // all to lower case
    s2 := []rune(s)
    var toLower []rune
    for _, char := range s2 {
        if char >= 'A' && char <= 'Z' {
            toLower = append(toLower, rune(char + 32))
        } else {
            toLower = append(toLower, rune(char))
        }
    }
    
    // -----------
    var toRevCap []rune

    for i, char := range toLower {
        if char == '9' {toRevCap = append(toRevCap, '9'); continue}
        if i + 1 < len(toLower) && toLower[i + 1] == ' ' || i == len(toLower) -1 {
            if char >= 'a' && char <= 'z' {
                toRevCap = append(toRevCap, rune(char - 32))
            } else {
                toRevCap = append(toRevCap, rune(char))
            }
        } else {
            toRevCap = append(toRevCap, rune(char))
        }
    }

    // -----------
    return string(toRevCap)
}

func println(s string) {
    for _, char := range s {
        z01.PrintRune(char)
    }
    z01.PrintRune('\n')
}


func main(){
    args := os.Args[1:]
    if len(args) == 0 {return}
    for _, arg := range args {
        println(reverseStrCap(arg))
    }
}