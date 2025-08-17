package piscine

func FifthAndSkip(str string) string {
    if str == "" {return "\n"}
    var result string
    var strToRunes []rune
    for _, char := range str{
        if char != ' '{
            strToRunes = append(strToRunes, char)
        }
    }
    if len(strToRunes) < 5 {return "Invalid Input\n"}
    i := 0
    for i + 5 < len(strToRunes) {
        result += string(strToRunes[i:i+5]) + " "
        i+=6
    }
    if i < len(str) {
        result += string(strToRunes[i:])
    }
    return result + "\n"
}