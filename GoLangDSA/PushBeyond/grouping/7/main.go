package main

import (
    "os"
    "fmt"
)

func Contains(text, toFind string) bool {
    for i := 0; i <= len(text)-len(toFind); i++ {
        if toFind == text[i:i+len(toFind)] {
            return true
        }
    }
    return false
}

func EndsWithSymbol(w string) bool {
    return (w[len(w)-1] >= 'a' && w[len(w)-1] <= 'z') || (w[len(w)-1] >= 'A' && w[len(w)-1] <= 'Z')
}

func Split(text, splitter string) (result []string) {
    var word string
    for _, char := range text {
        if string(char) != splitter {
            word += string(char)
        } else if word != "" {
            result = append(result, word)
            word = ""
        }
    }
    if word != "" {
        result = append(result, word)
    }
    return
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

func Print(text []string) {
    for i, str := range text {
        if !EndsWithSymbol(str) {str = str[:len(str)-1]}
        fmt.Printf("%d: %s\n", i+1, str)
    }
}

func main(){
    args := os.Args[1:]
    
    if len(args) != 2 {return}
    exp := args[0]
    
    if exp[0] != '(' || exp[len(exp)-1] != ')' {return}
    exp = exp[1:len(exp)-1]

    var expToSlice []string

    if Contains(exp, "|") {
        expToSlice = append(expToSlice, Split(exp, "|")...)
    } else {
        expToSlice = append(expToSlice, exp)
    }

    text := args[1]
    var textToSlice []string
    textToSlice = Split(text, " ")

    getResult := Search(expToSlice, textToSlice)
    Print(getResult)
}



