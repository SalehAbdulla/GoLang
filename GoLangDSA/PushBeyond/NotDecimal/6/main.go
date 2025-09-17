package piscine

import (
    "strconv"
)

// If the number doesn't have a decimal point or 
// there is only a zero after the . 
// return the number followed by a newline \n.
// ----
// If the argument is empty return a newline \n.
// If the argument is not a number return it followed by a newline \n.


func NotDecimal(dec string) string {
    decIndex := -1

    for i, char := range dec {
        if char == '.' {
            decIndex = i
            break
        }
    }

    if decIndex == -1 {return dec + "\n"}
    
    onlyZerosAfterDecPoint := true
    for i := decIndex; i < len(dec); i++ {
        if dec[i] != '0' {
            onlyZerosAfterDecPoint = false
            break
        }
    }
    
    if onlyZerosAfterDecPoint {return "\n"}

    n, ok := Atoi(dec)
    if !ok {return dec + "\n"}

    return strconv.Itoa(n) + "\n"
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
        if char == '.' {continue}
        if !(char >= '0' && char <= '9') {
            return -1, false
        } else {
            result *= 10
            result = int(char - '0') + result
        }
    }
    if isNegative {
        return result *-1, true
    }
    return result, true
}



