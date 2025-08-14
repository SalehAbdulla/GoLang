package piscine

func ItoaBase(value, base int) string {
    if value == 0 {return "0"}

    digits := "0123456789ABCDEF"
    isNegative := value < 0
    var uValue uint64
    if isNegative {
        uValue = uint64(-int64(value))
    } else {
        uValue = uint64(value)
    }
    uBase := uint64(base)

    var result string
    for uValue > 0 {
        reminder := uValue % uBase
        result = string(digits[reminder]) + result
        uValue /= uBase
    }
    if isNegative{
        return "-" + result
    }
    return result
}
