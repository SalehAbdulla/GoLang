package piscine

func PrintIf(str string) string {
	//If it's an empty string return G followed by a newline \n.
	if str == "" {
		return "G" + "\n"
	}
	//	returns the letter G followed by a newline \n if the argument length is more or equal than 3, otherwise returns Invalid Input followed by a newline \n.
	if len(str) >= 3 {
		return "G" + "\n"
	} else {
		return "Invalid Input" + "\n"
	}

}
