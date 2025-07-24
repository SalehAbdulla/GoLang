package piscine

func RepeatAlpha(s string) string {
    result := ""
    for _, char := range s {
        if char >= 'a' && char <= 'z' {
            getAsciiVal := int(char - 96) // 1->a
            for i:= 0; i < getAsciiVal; i++ {
                result += string(char)
            }
        } else if char >= 'A' && char <= 'Z' {
            getAsciiVal := int(char - 64) // 1 -> A
            for i := 0; i < getAsciiVal; i++ {
                result += string(char)
            }
        }  else {
            result += string(char)
        }
    }
    return result
}
