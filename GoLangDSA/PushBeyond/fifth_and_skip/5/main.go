package piscine

import (
    "strings"
)

func FifthAndSkip(str string) (result string) {
    if str == "" {return "\n"}
    str = strings.ReplaceAll(str, " ", "")
    if len(str) < 5 {return "Invalid Input\n"}
    strToRune := []rune(str)
    for i, char := range strToRune {
        if (i+1)%6 == 0 {
            result += " "
            continue
        } else {
            result += string(char)
        }
    }
    result += "\n"
    return
}