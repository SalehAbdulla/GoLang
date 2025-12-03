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


func endsWithPunc(s string) bool {
    if !(s[len(s)-1] >= 'a' && s[len(s)-1] <= 'z') && !(s[len(s)-1] >= 'A' && s[len(s)-1] <= 'Z') {
        return true
    }
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

func printResult(result []string) {
    for i, res := range result {
        if endsWithPunc(res) {
            res = res[:len(res) - 1]
            fmt.Printf("%d: %s\n", i+1, res)
        } else {
            fmt.Printf("%d: %s\n",i+1, res)
        }
    }
}

func main(){

    // if the number of arguments is different from 2, if the regular expression is not valid, 
    // if the last argument is empty or if there are no matches, the program should print nothing.

    args := os.Args[1:]
    if len(args) != 2 { return }

    regex := args[0]
    text := args[1]

    // if the regular expression is not valid 
    if len(regex) < 3 || regex[0] != '(' || regex[len(regex) - 1] != ')' { return }
    regex = regex[1:len(regex)-1] // remove prackets

    // turn exp into []string + handle | operator
    var regexSplit []string
    if contains(regex, "|") {
        regexSplit = split(regex, "|")
    } else {
        regexSplit = append(regexSplit, regex)
    }

    // if text is empty
    if text == "" { return }

    var textSplit []string
    textSplit = split(text, " ")

    result := search(regexSplit, textSplit)
    printResult(result)
}