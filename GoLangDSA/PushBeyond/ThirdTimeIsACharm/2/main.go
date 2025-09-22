package piscine

// Return the output followed by a newline \n.
// If the string is empty, return a newline \n.
// If there is no third character, return a newline \n.

func ThirdTimeIsACharm(str string) string {
    if str == "" || len(str) < 3 {return "\n"}
    var result string

    for i := 2; i < len(str); i += 3 {
        result += string(str[i])
    }


    return result + "\n"
}