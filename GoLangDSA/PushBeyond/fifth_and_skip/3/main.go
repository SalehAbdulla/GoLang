package piscine

func FifthAndSkip(str string) string {
    if str == "" {return "\n"}

    strToRunes := removeSpacesSlice(str)

    if len(strToRunes) < 5 {return "Invalid Input\n"}

    var result string
    i := 0
    for i + 5 < len(strToRunes) {
        result += string(strToRunes[i:i+5]) + " "
        i += 6
    }
    if i != len(strToRunes) {
        result += string(strToRunes[i:])
    }
    return result + "\n"
}

func removeSpacesSlice(str string) (strToRunes []rune) {
    for _, char := range str {
        if char != ' ' {
            strToRunes = append(strToRunes, char)
        }
    }
    return
}
