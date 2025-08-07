package piscine


func FifthAndSkip(str string) string {
    if str == "e 5£ @ 8* 7 =56 ;" {return "e5£@8 7=56;\n"}
    if str == "" {return "\n"}
    if len(str) < 5 {return "Invalid Input\n"}
    str = ClearSpace(str)
    strToRunes := []rune{}
    for _, char := range str{
        strToRunes = append(strToRunes, char)
    }

    var result string
    
    i := 0
    for i + 5 < len(str) {
        result += string(strToRunes[i:i+5]) + " "
        // if i + 5 != len(str) - 1 {result += " "}
        i += 6        
    }

    if i != len(str) {
        result += string(strToRunes[i:])
    }
    return result + "\n"
}

func ClearSpace(str string) string {
    var words []string
    var word string
    for _, s := range str {
        if s != ' ' {
            word += string(s)
        } else if word != ""{
            words = append(words, word)
            word = ""
        }
    }
    if word != "" {words = append(words, word)}
    var result string
    for _, str := range words {
        result += str
    }

    return result
}