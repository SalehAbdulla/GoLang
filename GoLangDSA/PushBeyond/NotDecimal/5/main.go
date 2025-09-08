package piscine

import (
    "strconv"
)

func NotDecimal(dec string) string {
    if dec == "" {return "\n"}

    decimalIndex := -1
    for i, char := range dec {
        if char == '.' {
            decimalIndex = i; break
        }
    }

    if decimalIndex == -1 {return dec + "\n"}
    var cleanDec string
    for _, char := range dec {
        if char != '.' {
            cleanDec += string(char)
        }
    }

    decToInt, ok := Atoi(cleanDec)
    if !ok {return dec + "\n"}

    return strconv.Itoa(decToInt) + "\n"
}

func Atoi(num string) (int, bool) {
    if num == "" {return -1, false}
    if num == "0" {return 0, true}

    isNegative := false
    if num[0] == '-' {
        isNegative = true
        num = num[1:]
    }
    var result int
    for _, char := range num {
        if char < '0' || char > '9' {
            return -1, false
        } else {
            result *= 10
            result += int(char - '0')
        }
    }

    if isNegative {
        return result * -1, true
    }
    return result, true

}