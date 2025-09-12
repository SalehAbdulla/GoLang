package piscine



func IsUpper(r rune) bool {
    if r >= 'A' && r <= 'Z' {return true}
    return false
}

func IsCamelCaseValid(s string) bool {
    for _, char := range s {
        if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' {
            continue
        } else {
            return false
        }
    }

    for i, char := range s {
        if IsUpper(char) && i == len(s)-1 {return false}
        if IsUpper(char) && i + 1 < len(s) && IsUpper(rune(s[i+1])) {return false}
    }
    return true
}

func CamelToSnakeCase(s string) string{
    if !IsCamelCaseValid(s) {return s}

    var result []rune
    for i, char := range s {
        if IsUpper(char) && i-1 >= 0  && i != 0 && i != len(s) - 1{
            result = append(result, '_') 
            result = append(result, char)
        } else {
            result = append(result, char)
        }
    }
    return string(result)
}
