package piscine

func SplitBySpace(s string) (words []string) {
    var word string
    for _, char := range s{
        if char != ' ' {
            word += string(char)
        } else if word != "" {
            words = append(words, word)
            word = ""
        }
    }
    if word != "" {words = append(words, word)}
    return
}


func WordFlip(str string) string{
    if str == "" {return "Invalid Output"}
    wordToSlice := SplitBySpace(str)
    var result string
    for i := len(wordToSlice) - 1; i >= 0; i-- {
        result += string(wordToSlice[i])
        if i != 0 {
            result += " "
        }
    }
    return result + "\n"
}