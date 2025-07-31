package piscine

func SaveAndMiss2(arg string, num int) string {
	if num <= 0 || arg == "" {
		return arg
	}

	result := ""
	save := true
	count := 0

	for _, char := range arg {
		if save {
			result += string(char)
		}
		count++
		if count == num {
			save = !save
			count = 0
		}
	}

	return result
}
