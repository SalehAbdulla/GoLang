package piscine


//	Write a function that takes a string as an argument and returns true if the string contains any number, otherwise return false.

func CheckNumber2(arg string) bool {
	isExist := false
	for _, char := range arg {
		if char >= '0' && char <= '9' {
			isExist = true
			return isExist
		}
	}
	return isExist
}
