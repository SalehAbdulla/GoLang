package piscine


func RetainFirstHalf(str string) string {
	// If the string is empty, return an empty string.
	// If the string length is equal to one, return the string.
	if str == "" || len(str) == 1{
		return str
	}
	
	// If the length of the string is odd, round it down.
	halfLength := len(str) / 2
	return str[:halfLength]
}
