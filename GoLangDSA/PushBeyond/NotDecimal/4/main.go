package piscine

import (
    "strconv"
)

func NotDecimal(dec string) string {
    if dec == "" {return "\n"}

    dotIndex := -1
    for i, c := range dec {
        if c == '.' {
            dotIndex = i; break
        }
    }

    if dotIndex == -1 {return dec + "\n"}

    allZerosAfterDot := true
    for i := dotIndex+1; i < len(dec); i++ {
        if dec[i] != '0' {allZerosAfterDot = false; break}
    }
    
    if allZerosAfterDot {return dec + "\n"}
    // ----------------
    var clearDec string
    for _, char := range dec {
        if char != '.' {
            clearDec += string(char)
        }
    }
    // ----------------
    toInt, passed := Atoi(clearDec)
    if !passed {return dec + "\n"}
    return strconv.Itoa(toInt) + "\n"

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
        return result *-1, true
    }

    return result, true
}   

