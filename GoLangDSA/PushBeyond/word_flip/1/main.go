package piscine

func WordFlip(str string) string {
    if str == "" {return "Invalid Output"}
    var result []string
    word := ""
    for _, char := range str {
        if char != ' ' {
            word += string(char)
        } else if word != "" {
            result = append(result, word)
            word = ""
        }
    }

    if word != "" {result = append(result, word)}

    if len(result) == 0 {return "\n"}



    var resultReverse []string
    for i := len(result)-1; i >= 0; i-- {
        resultReverse = append(resultReverse, result[i])
    }

    var finalResult string

    for i, text := range resultReverse {
        for _, char := range text {
            finalResult += string(char)
        }
        if i != len(resultReverse) -1 {
            finalResult += " "
        } else {
            finalResult += "\n"
        }
    }
    return finalResult
}
