package piscine

func split(text string) []string {
    var words []string
    var word string
    for _, char := range text {
        if char != ' ' {
            word += string(char)
        } else if word != "" {
            words = append(words, word)
            word = ""
        }
    }
    if word != "" {words = append(words, word)}
    return words
}

func WordFlip(str string) string {
    if str == "" {return ""}
    strSlice := split(str)

    var result string
    for i := len(strSlice) - 1; i >= 0; i-- {
        result += string(strSlice[i])
        if i != 0 {result += " "}
    }

    return result + "\n"
}
