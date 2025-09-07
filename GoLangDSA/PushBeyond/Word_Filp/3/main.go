package piscine

import (
    "strings"
)

func WordFlip(str string) string {
    if str == "" {return str}
    str = strings.TrimSpace(str)
    strToSlice := strings.Split(str, " ")
    var result string
    for i := range strToSlice {
        word := strings.TrimSpace(strToSlice[len(strToSlice)-1-i])
        if word == "" {continue}
        result += word
        if i != len(strToSlice) - 1 {result += " "}
    }
    return result + "\n"
}
