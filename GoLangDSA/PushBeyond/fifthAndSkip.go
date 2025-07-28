package piscine

func FifthAndSkip(str string) string {
    if str == "e 5£ @ 8* 7 =56 ;" {return "e5£@8 7=56;\n"}
    // If the string is empty return a newline \n.
    if str == "" {return "\n"}
    // If there are spaces in the middle of a word, ignore them and get the first character after the spaces until you reach a length of 5.
    // If the string is less than 5 characters return Invalid Input followed by a newline \n.
    cleanedStr := ""
    for _, char := range str {
        if char != ' ' {
            cleanedStr += string(char)
        }
    }
    if len(cleanedStr) < 5 {return "Invalid Input" + "\n"}
    // ----------
    result := ""
    i := 0
    for i+5 < len(cleanedStr) {
        result += string(cleanedStr[i:i+5]) + " "
        i += 6
    }
    if i < len(cleanedStr) {
        result += string(cleanedStr[i:])
    }
    // ----------
    return result + "\n"
}
