package main
import (
	"fmt"
	"piscine"
)

//Command failed: go run ./tests/fifthandskip_test/main.go
// FifthAndSkip("1234556789") == "12345 67891 34556 89\n" instead of "12345 6789\n"
// exit status 1


func main() {
	fmt.Print(piscine.FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	fmt.Print(piscine.FifthAndSkip("This is a short sentence"))
	fmt.Print(piscine.FifthAndSkip("1234556789"))
}

func FifthAndSkip(str string) string {
	// If the string is empty return a newline \n.
	if str == "" {return "\n"}
	// After cleaning -- If the string is less than 5 characters return Invalid Input followed by a newline \n.
	cleanStr := []rune(str)
	for _, char := range str {
		if char != ' ' {
			cleanStr = append(cleanStr, char) 
		}
	}
	if len(cleanStr) < 5 {return "Invalid Input\n"}
	/// ----------------
	result := ""
	i := 0
	for i+5 < len(cleanStr) {
		result += string(cleanStr[i:i+5]) + " "
		i += 6
	}
	if i != len(cleanStr) {
		result += string(cleanStr[i:])
	}

	/// ----------------
	return result + "\n"
}
