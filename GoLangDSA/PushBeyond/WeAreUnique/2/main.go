package piscine

func WeAreUnique(str1 , str2 string) int {
    if str1 == "" && str2 == "" {return -1}
    counter := 0
    hashMap1 := make(map[rune]bool)
    hashMap2 := make(map[rune]bool)
    used := make(map[rune]bool)

    for _, ch := range str1 {hashMap1[ch] = true}
    for _, ch := range str2 {hashMap2[ch] = true}

    for _, ch := range str1 {
        if !hashMap2[ch] && !used[ch] {used[ch] = true; counter++}
    }

    for _, ch := range str2 {
        if !hashMap1[ch] && !used[ch] {used[ch] = true; counter++}
    }

    return counter
}
