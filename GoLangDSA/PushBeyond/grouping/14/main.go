package main

import (
    "os"
    "fmt"
)

func contains(text, toFind string) bool {
    for i := 0; i <= len(text)-len(toFind); i++ {
        if toFind == text[i:i+len(toFind)] {return true}
    }
    return false
}

func split(text, splitter string) []string {
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

func printNeeded(strs []string) {
    for i, str := range strs {
        if !endsWithChar(str) {str = str[:len(str)-1]}
        fmt.Printf("%d: %s\n", i, str)
    }
}

func endsWithChar(str string)bool {
    l := len(str)
    return (str[l] >= 'a' && str[l] <= 'z') || (str[l] >= 'A' && str[l] <= 'Z') 
}

func search(exps, texts []string) (result []string) {
    for _, text := range texts {
        for _, exp := range exps {
            if contains(text, exp) {
                result = append(result, text)
            }
        }
    }
    return
}

func main(){
    args := os.Args[1:]
    if len(args) != 2 {return}
    regexp := args[0]
    
    if len(regexp) < 2 || regexp[0] != '[' || regexp[len(regexp)-1] != ']' {return}
    regexp = regexp[1:len(regexp)-1] // pop

    var regexpSlice []string
    if contains(regexp, "|") {
        regexpSlice = append(regexpSlice, split(regexp, "|")...)
    } else {
        regexpSlice = append(regexpSlice, regexp)
    }

    text := args[1]
    if text == "" {return}
    var textSlice []string
    textSlice = append(textSlice, split(text, " ")...)
    getResult := search(regexpSlice, textSlice)
    printNeeded(getResult)
}
