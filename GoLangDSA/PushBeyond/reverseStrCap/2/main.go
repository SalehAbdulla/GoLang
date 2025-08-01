package main

import (
    "os"
    "github.com/01-edu/z01"
)

func main(){
    args := os.Args[1:]
    if len(args) == 0 {return}
    
    for _, arg := range args {
        toLower := toLowerCase(arg)
        revCap := reverseCap(toLower)
        print(revCap)
        z01.PrintRune('\n')
    }
}

// --- toLowerCase
func toLowerCase(s string) string {
    var result string
    for _, char := range s {
        if char >= 'A' && char <= 'Z' {
            result += string(rune(char + 32))
            continue
        } else {
            result += string(char)
        }
    }
    return result
}

// --- reverseCap
func reverseCap(s string) string {
    var result string

    for i := 0; i < len(s); i++ {
        if i + 1 < len(s) && s[i+1] == ' ' || i == len(s) - 1 {
            if s[i] >= 'a' && s[i] <= 'z' {
                result += string(rune(s[i]-32))
                continue
            }
        }
        result += string(s[i])
    }
    return result
}

// --- print
func print(s string) {
    for _, c := range s {
        z01.PrintRune(c)
    }
}
