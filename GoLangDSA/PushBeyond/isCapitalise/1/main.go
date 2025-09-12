package piscine

func isWithinSmallChar(r rune) bool {
    if r >= 'a' && r <= 'z' {return true}
    return false
}

func IsCapitalized(s string) bool {
    if len(s) == 0 {return false}
    for i, char := range s {
        if isWithinSmallChar(char) && i == 0 {return false}
        if char == ' ' && i+1 < len(s) && isWithinSmallChar(rune(s[i+1])) {return false}
    }
    return true
}
