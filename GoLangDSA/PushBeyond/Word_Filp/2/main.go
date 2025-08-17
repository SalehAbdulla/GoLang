package piscine

func textToSlice(s string) (result []string) {
    var word string
    for _, char := range s{
        if char != ' ' {
            word += string(char)
        } else if word != "" {
            result = append(result, word)
            word = ""
        }
    }
    if word != "" {result = append(result, word)}
    return
}

func WordFlip(str string) string {
    var result string
    strToSlice := textToSlice(str)
    for i, _ := range strToSlice {
        result += strToSlice[len(strToSlice)-1-i]
        if i != len(strToSlice)-1{
            result += " "
        }
    }
    return result + "\n"
}
