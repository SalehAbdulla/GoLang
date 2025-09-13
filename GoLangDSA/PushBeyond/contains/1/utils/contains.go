package utils

func contains(s, toFind string) bool {
    for i := 0; i <= len(s) - len(toFind); i++ {
        if toFind == s[i:i+len(toFind)] {return true}
    }
    return false
}