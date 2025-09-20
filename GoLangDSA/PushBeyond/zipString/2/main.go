package piscine

import (
    "strconv"

)

func ZipString(s string) string {
    var result string
    counter := 1
    for i, char := range s {
        if i + 1 < len(s) && char == rune(s[i+1]) {
            counter++
        } else {
            result += strconv.Itoa(counter)
            result += string(char)
            counter = 1
        }
    }
    return result
}