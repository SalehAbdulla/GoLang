package piscine


const digits = "0123456789ABCDEF"
func ItoaBase(value, base int) string {
    if value == 0 {return "0"}
    var uValue uint64

    isNegative := value < 0
    if isNegative {
        uValue = uint64(-int64(value))
    } else {
        uValue = uint64(value)
    }

    var result string

    for uValue > 0 {
        remiander := uValue % uint64(base)
        result = string(digits[remiander]) + result
        uValue /= uint64(base)
    }

    if isNegative{
        return "-" + result
    } else {
        return result
    }
}