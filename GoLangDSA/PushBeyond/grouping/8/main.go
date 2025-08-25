package main

import (
    "os"
    "fmt"
)


func Contains(text, toFind string) bool {
    for i := 0; i <= len(text)-len(toFind); i++ {
        if toFind == text[i:i+len(toFind)] {return true}
    }
    return false
}
func Split(text, splitter string) []string {
    var words []string
    var word string 
    for _, char := range text {
        if string(char) != splitter {
            word += string(char)
        } else if word != "" {
            words = append(words, word)
            word = ""
        }
    }
    if word != "" {words = append(words, word)}
    return words
}

func Search(exps, texts []string) (result []string) {
    for _, text := range texts {
        for _, exp := range exps {
            if Contains(text, exp) {
                result = append(result, text)
            }
        }
    }
    return
}

func EndsWithCharacter(t string) bool {
    return (t[len(t)-1] >= 'a' && t[len(t)-1] <= 'z') || (t[len(t)-1] >= 'A' && t[len(t)-1] <= 'Z')
}

func PrintAsNeeded(words []string) {
    for i, word := range words {
        if !EndsWithCharacter(word) {word=word[:len(word)-1]}
        fmt.Printf("%d: %s\n", i+1, word)
    }
}

func main() {
    args := os.Args[1:]
    if len(args) != 2 || args[1] == "" {return}
    
    regExp := args[0]
    if regExp[0] != '(' || regExp[len(regExp)-1] != ')' {return}

    regExp = regExp[1:len(regExp)-1]

    var regExpToSlice []string
    if Contains(regExp, "|") {
        regExpToSlice = append(regExpToSlice, Split(regExp, "|")...)
    } else {
        regExpToSlice = append(regExpToSlice, regExp)
    }

    text := args[1]
    var textToSlice []string
    textToSlice = append(textToSlice, Split(text, " ")...)
    getSearchRes := Search(regExpToSlice, textToSlice)
    PrintAsNeeded(getSearchRes)

}

