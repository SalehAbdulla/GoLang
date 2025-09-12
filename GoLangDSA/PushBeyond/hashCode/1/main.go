package piscine

// (ASCII of current character + size of the string) % 127,
func HashCode(dec string) string {
    var result string
    for _, char := range dec {
        encoded := (int(char) + len(dec) ) % 127
        if encoded < 33 {encoded += 33}
        result += string(rune(encoded))
    }
    return result
}