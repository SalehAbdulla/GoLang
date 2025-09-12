package piscine

// WeAreUnique("everyone", "") == 8 instead of 6
func WeAreUnique(str1, str2 string) int {
	if str1 == "" && str2 == "" {
		return -1
	}
	// if str1 == "" && str2 != "" {return len(str2)}
	// if str1 != "" && str2 == "" {return len(str1)}
	hashMap1 := make(map[rune]bool)
	hashMap2 := make(map[rune]bool)
	used := make(map[rune]bool)
	for _, char := range str1 {
		hashMap1[char] = true
	}
	for _, char := range str2 {
		hashMap2[char] = true
	}
	var counter int
	for _, char := range str1 {
		if !hashMap2[char] && !used[char] {
			counter++
			used[char] = true
		}
	}
	for _, char := range str2 {
		if !hashMap1[char] && !used[char] {
			counter++
			used[char] = true
		}
	}
	return counter
}
