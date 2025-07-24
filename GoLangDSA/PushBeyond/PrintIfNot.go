package piscine

//Write a function that takes a string as an argument and 
// returns the letter G if the argument length is less than 3, otherwise returns Invalid Input followed by a newline \n.
// If it's an empty string return G followed by a newline \n.


func PrintIfNot(str string) string {

	if len(str) < 3 || str == "" {
		return "G" + "\n"
	} else {
		return "Invalid Input" + "\n"
	}
}
