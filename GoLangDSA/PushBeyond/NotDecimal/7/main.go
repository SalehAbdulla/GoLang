package piscine

import (
    "strconv"
)

func NotDecimal(dec string) string {
    if dec == "0" {return "0"}

    decimalIndex := -1
    for i, char := range dec {
        if char == '.' {decimalIndex = i; break}
    }

    if decimalIndex == -1 {return dec + "\n"}
    allzerosAfterDot := true
    for i := decimalIndex; i < len(dec); i++ {
        char := dec[i]
        if char != '0' {allzerosAfterDot = false; break}
    }

    if allzerosAfterDot {return dec + "\n"}

    var clearStr string
    for _, char := range dec {
        if char != '.' {
            clearStr += string(char)
        }
    }

    strToInt, ok := Atoi(clearStr)
    if !ok {return dec + "\n"}
    return strconv.Itoa(strToInt) + "\n"


}


func Atoi(num string) (int, bool) {
    if num == "0" {return 0, true}

    isNegative := false
    if num[0] == '-' {
        isNegative = true
        num = num[1:]
    }

    var result int
    for _, char := range num {
        if char < '0' || char > '9' {return -1, false}
        result *= 10
        result += int(char - '0')
    }

    if isNegative {
        return result * -1, true
    }
    return result, true
}