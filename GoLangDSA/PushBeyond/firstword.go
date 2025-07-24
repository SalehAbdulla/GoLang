package piscine

func FirstWord(s string) string {
    i := 0
    for i < len(s) && s[i] == ' ' {
        i++
    }
    start := i

    // ------

    for i < len(s) && s[i] != ' ' {
        i++
    }
    ends := i

    return s[start:ends]
}
