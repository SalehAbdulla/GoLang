package piscine

func ItoaBase(value, base int) (result string) {
    digits := "0123456789ABCDEF"
    var uValue uint64
    isNegative := value < 0
    if isNegative {
        uValue = uint64(-int64(value))
    } else {
        uValue = uint64(value)
    }
    uBase := uint64(base)
    for uValue > 0 {
        getReminder := uValue % uBase
        result = string(digits[getReminder]) + result
        uValue /= uBase
    }
    if isNegative {
        return "-" + result
    }
    return result
}
