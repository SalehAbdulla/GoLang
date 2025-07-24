package piscine

//Write a function that takes a string as an argument and 
// returns true if the string contains any number, otherwise return false.
func CheckNumber(arg string) bool {
	// numberExist
	numberExist := false
	for _, char := range arg {
		if int(char) >= 0 && int(char) - '0' <= 9 {
			numberExist = true
		}
	}
	return numberExist
}
