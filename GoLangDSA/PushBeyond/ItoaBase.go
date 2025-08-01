package piscine


func ItoaBase(value, base int) string {
    if value == 0 {return "0"} // all bases zero is zero
    const digits = "0123456789ABCDEF" // Declare Base Constant 
    var uvalue uint64 

    isNegative := value < 0
    // if Negative, safely convert it into positive uint64(-int64(uvalue))
    if isNegative {
        uvalue = uint64(-int64(value))
    } else {
        uvalue = uint64(value)
    }
    // --------------------------
    result := ""
    // ItoaBase steps
    // 1) take the reminder uvalue & unint(base)
    // 2) append(digit[reminder])
    // 3) Divide uvalue/base

    for uvalue > 0 {    
        reminder := uvalue % uint64(base)
        result = string(digits[reminder]) + result
        uvalue /= uint64(base)
    }

    // append "-" if negative
    if isNegative {
        return "-" + result
    }
    return result

}