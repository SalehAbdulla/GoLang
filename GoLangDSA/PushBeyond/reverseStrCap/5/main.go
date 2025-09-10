package main
//                  |
// $ go run . "First SMALL TesT" | cat -e
// firsT smalL tesT$

import (
    "os"
    "github.com/01-edu/z01"
)

func ToLower(s string) string {
    var result string
    for _, r := range s {
        if r >= 'A' && r <= 'Z' {
            result += string(rune(r + 32))
            continue
        } else {
            result += string(r)
        }
    }
    return result
}

func ToUpper(r rune) (result rune) {
    if r >= 'a' && r <= 'z' {
        return rune(r - 32)
    } else {
        return r
    }
}

func ToRevCap(s string) string {
    sToSlice := []rune(s)
    for i := 0; i < len(sToSlice); i++ {
        r := sToSlice[i]
        if r == ' ' && i-1 >= 0 && sToSlice[i-1] >= 'a' && sToSlice[i-1] <= 'z' {
            prevChar := sToSlice[i-1]
            sToSlice[i-1] = ToUpper(rune(prevChar))
        } else if i == len(sToSlice) - 1 && sToSlice[i] >= 'a' && sToSlice[i] <= 'z' {
            lastChar := sToSlice[i]
            sToSlice[i] = ToUpper(rune(lastChar))
        } 
    }
    return string(sToSlice)
}

func Print(s string) {
    for _, r := range s {
        z01.PrintRune(r)
    }
}

func main() {
    args := os.Args[1:]
    if len(args) < 1 {return}
    for _, arg := range args {
        arg = ToLower(arg)
        Print(ToRevCap(arg))
        Print("\n")
    }
}
