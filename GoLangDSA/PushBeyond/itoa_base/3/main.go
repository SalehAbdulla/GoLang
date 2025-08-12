package piscine

func ItoaBase(value, base int) string {
    if value == 0 {return "0"}
    digits := "0123456789ABCDEF"

    var uValue uint32
    isNegative := value < 0
    if isNegative {
        uValue = uint32(-int32(value))
    } else {
        uValue = uint32(value)
    }

    var result string
    
    for uValue > 0 {
        reminder := uValue % uint32(base)
        result = string(digits[reminder]) + result
        uValue /= uint32(base)
    }

    if isNegative {
        return "-" + result
    }
    return result

}