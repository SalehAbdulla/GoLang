package piscine

func IsChar(r byte) bool {
    return r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z'
}

func IsLower(r byte) bool {
    return r >= 'a' && r <= 'z'
}

func IsCapitalized(s string) bool {
    if s == "" {return false}
    if s[0] >= 'a' && s[0] <= 'z'{return false}
    for i := 1; i < len(s); i++ {
        curr := s[i]
        before := s[i-1]
        if IsChar(curr) && before == ' ' && IsLower(curr) {
            return false
        }
    }
    return true
}
