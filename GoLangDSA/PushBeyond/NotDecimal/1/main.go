package piscine 

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
    isAllZerosAfterDot := true
    for i := dotIndex + 1; i < len(dec); i++ {
        if dec[i] != '0' {
            isAllZerosAfterDot = false
            break
        }
    }

    if isAllZerosAfterDot {return dec + "\n"}
    // ------------------
    var clearFromDot string
    for _, char := range dec {
        if char == '.' {
            continue
        } else {
            clearFromDot += string(char)
        }
    }
    decToInt, ok := Atoi(clearFromDot)
    if !ok {return dec + "\n"}
    decToStr := Itoa(decToInt)
    return decToStr + "\n"

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
