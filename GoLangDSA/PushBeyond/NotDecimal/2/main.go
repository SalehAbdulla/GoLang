package piscine

// import (
//     "strconv"
// )


// If the argument is not a number return it followed by a newline \n.



func NotDecimal(dec string) string {
    if dec == "" {return "\n"}
    dotIndex := -1
    for i, char := range dec {
        if char == '.' {
            dotIndex = i
            break
        }
    }

    if dotIndex == -1 {return dec + "\n"}
    zerosAfterDot := true
    for i := dotIndex + 1; i < len(dec); i++ {
        if dec[i] != '0' {
            zerosAfterDot = false
            break
        }
    }
    if zerosAfterDot {return dec + "\n"}

    /// -------------- 
    var decCleaned string
    for _, char := range dec {
        if char != '.' {
            decCleaned += string(char)
        }
    }

    /// --------------
    decTonum, ok := Atoi(decCleaned)
    if !ok {return dec + "\n"}
    return Itoa(decTonum) + "\n"
}


func Itoa(num int) string {
    if num == 0 {return "0"}
    isNegative := false

    if num < 0 {
        isNegative = true
        num = -num
    }

    var result string

    for num > 0 {
        digit := num % 10
        result = string(rune(digit + '0')) + result
        num /= 10
    }

    if isNegative {
        return "-" + result
    }
    return result

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
    for _, n := range num {
        if n < '0' || n > '9' {
            return -1, false
        } else {
            result *= 10
            result = int(n - '0') + result
        }
    }

    if isNegative {
        return result * -1, true
    } else {
        return result, true
    }
}
