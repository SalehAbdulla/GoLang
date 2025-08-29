package piscine

import (
    "strings"
)

func FifthAndSkip(str string) string {
    if str == "" {return "\n"}
    // --------
    str = strings.ReplaceAll(str, " ", "")
    if len(str) < 5 {return "Invalid Input\n"}
    // --------
    strToRune := []rune(str)
    var result string

    i := 0
    for i + 5 < len(strToRune) {
        result += string(strToRune[i:i+5]) + " "
        i += 6 
    }

    if i != len(strToRune) {
        result += string(strToRune[i:])
    }

    return result + "\n"
}
