package main

import (
    "fmt"
    "os"
)

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

func contains(text, toFind string) bool {
    for i := 0; i <= len(text) - len(toFind); i++ {
        if toFind == text[i:i+len(toFind)] {return true}
    }
    return false
}

func endsWithChar(word string) bool {
    if len(word) == 0 {return false}
    r := word[len(word)-1]
    if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' {return true}
    return false
}

func search(exps, texts []string) []string {
    var result []string
    for _, text := range texts {
        for _, exp := range exps {
            if contains(text, exp) {
                result = append(result, text)
            }
        }
    }
    return result
}

func printAsNeeded(strs []string) {
    for i, str := range strs {
        if !endsWithChar(str) {str = str[:len(str)-1]} // pop
        fmt.Printf("%d: %s\n", i+1, str)
    }
}


func main(){
    args := os.Args[1:]
    if len(args) != 2 {return}
    exp := args[0]
    if len(exp) < 3 || exp[0] != '(' || exp[len(exp)-1] != ')' {return}

    exp = exp[1:len(exp)-1]
    var expToSlice []string
    if contains(exp, "|") {
        expToSlice = append(expToSlice, split(exp, "|")...)
    } else {
        expToSlice = append(expToSlice, exp)
    }

    text := args[1]
    if text == "" {return}
    var textToSlice []string
    textToSlice = append(textToSlice, split(text, " ")...)

    getResult := search(expToSlice, textToSlice)
    printAsNeeded(getResult)
}
