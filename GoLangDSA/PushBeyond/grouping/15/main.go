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

func printArr(strs []string) {
    for i, str := range strs {
        if !endsWithChar(str) {str = str[:len(str)-1]}
        fmt.Printf("%d: %s\n", i+1, str)
    }
}

func endsWithChar(word string) bool {
    l := len(word) - 1
    return word[l] >= 'a' && word[l] <= 'z' || word[l] >= 'A' && word[l] <= 'Z'
}


func main(){
    args := os.Args[1:]
    if len(args) != 2 {return}

    exp := args[0]
    if len(exp) < 3 || exp[0] != '(' || exp[len(exp)-1] != ')' {return}

    exp = exp[1:len(exp)-1] // pop

    var exps []string
    if contains(exp, "|") {
        exps = append(exps, split(exp, "|")...)
    } else {
        exps = append(exps, exp)
    }
    
    text := args[1]
    if text == "" {return}
    var texts []string
    texts = append(texts, split(text, " ")...)

    getResult := search(exps, texts)
    printArr(getResult)
}
