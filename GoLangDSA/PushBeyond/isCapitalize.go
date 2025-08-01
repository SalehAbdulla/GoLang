package piscine

func IsCapitalized(s string) bool {
    if s == "" {return false}
    if s[0] >= 'a' && s[0] <= 'z' {return false}
    for i := 0; i < len(s); i++ {
        if rune(s[i]) == ' ' {
            if i + 1 < len(s) && s[i+1] >= 'a' && s[i+1] <= 'z' {
                return false
            }
        }
    }
    return true
}
