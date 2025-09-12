package piscine

func Itoa(n int) string {
    if n == 0 {return "0"}
    isNegative := false
    if n < 0 {
        isNegative = true
        n = -n
    }

    var result string
    for n > 0 {
        digit := n % 10
        result = string(rune(int(digit + '0'))) + result
        n /= 10
    }

    if isNegative {
        return "-" + result
    }
    return result
}
