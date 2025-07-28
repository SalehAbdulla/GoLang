package piscine

func FifthAndSkip(str string) string {
    // If the string is empty return a newline \n.
    if str == "" {return "\n"}
    cleanedStr := []rune{}
    for _, char := range str {
        if char != ' '{
            cleanedStr = append(cleanedStr, char)
        }
    }
    if len(cleanedStr) < 5 {return "Invalid Input\n"}
    // --------------------------------
    result := ""
    i := 0
    for i + 5 < len(cleanedStr) {
        result += string(cleanedStr[i:i+5]) + " "
        i += 6
    }

    if i != len(cleanedStr) {
        result += string(cleanedStr[i:])
    }

    return result + "\n"
}
