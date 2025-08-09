package piscine

func SaveAndMiss3(arg string, num int) string {
	var result string
	saveMiss := true
	i := 0
	for i+num < len(arg) {
		if saveMiss {
			result += arg[i : i+num]
		}
		saveMiss = !saveMiss
		i += num
	}
	if saveMiss {
		result += string(arg[i:])
	}
	return string(result)
}
