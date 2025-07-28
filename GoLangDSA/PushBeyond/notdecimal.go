package piscine

func NotDecimal(dec string) string {
    // If the argument is empty return a newline \n.
    if dec == "" {return "\n"}

    // If the number doesn't have a decimal point or there is only a zero after the . 
    dotIndex := -1
    for i, char := range dec {
        if char == '.' {
            dotIndex = i
            break
        }
    }
    if dotIndex == -1 {return dec + "\n"}
    // ----
    //or there is only a zero after the .
    isAllZeros := true
    for i := dotIndex+1; i < len(dec); i++ {
        if dec[i] != '0' {
            isAllZeros = false
            break
        }
    }
    if isAllZeros {return dec + "\n"}
    // ---------- Smart way to remove trailing zeros
    // Clean zeros
    clear := ""
    for _, char := range dec {
        if char != '.' {
            clear += string(char)
        }
    }
    // -------------- 
    num := Atoi(clear)
    if num == -1 {return dec + "\n"}
    return Itoa(num) + "\n"
}



func Atoi(number string) int {
    if number == "0" {return 0}
    if number == "" {return -1}
    isNegative := false
    if number[0] == '-' {
        isNegative = true
        number = number[1:]
    }
    /// ---------------
    result := 0
    for _, char := range number {
        if char >= '0' && char <= '9' {
            result *= 10
            result += int(char - '0')
        } else {
            return -1
        }
    }
    /// ---------------
    if isNegative {
        return result * -1
    } else {
        return result
    }

}



func Itoa(number int) string {
    if number == 0 {return "0"}
    isNegative := false
    if number < 0 {
        isNegative = true
        number = -number
    }
    /// ---------------
    result := ""
    for number > 0 {
        digit := number % 10
        result = string(rune(digit + '0')) + result
        number /= 10
    }
    /// ---------------
    if isNegative {
        return "-" + result
    } else {
        return result 
    }

}