package main

import (
    "fmt"
    "os"
)

func contains(text, toFind string) bool {
    for i := 0; i <= len(text) - len(toFind); i++ {
        if toFind == string(text[i:i+len(toFind)]) {return true}
    }
    return false
}

func split(text, splitter string) (result []string) {
    var word string
    for _, char := range text {
        if string(char) != splitter {
            word += string(char)
        } else if word != "" {
            result = append(result, word)
            word = ""
        }
    }
    if word != "" {result = append(result, word)}
    return
}

func searchInText(search, text []string) (result []string) {
    for _, t := range text {
        for _, s := range search {
            if contains(t, s) {
                result = append(result, t)
            }
        }
    }
    return
}

func endsWithSymbol(s string) bool {
    return  !(s[len(s)-1] >= 'a' && s[len(s)-1] <= 'z') && !(s[len(s)-1] >= 'A' && s[len(s)-1] <= 'Z')
}

func removeEmpty(s []string) (result []string) {
    for _, str := range s {
        if str == "" || str == " " {
            continue
        } else {
            result = append(result, string(str))
        }
    }
    return
}

func main(){
    args := os.Args[1:]
    if len(args) != 2 {return}
    exp := args[0]
    textToSearchIn := args[1]
    if exp[0] != '(' || exp[len(exp)-1] != ')' {return}
    exp = exp[1:len(exp)-1]

    var cleanExp []string 

    if contains(exp, "|") {
        cleanExp = split(exp, "|")
    } else {
        cleanExp = append(cleanExp, exp) // Remove Brackets
    }

    getResult := searchInText(removeEmpty(cleanExp), removeEmpty(split(textToSearchIn, " ")))
    for i, word := range getResult {
        if endsWithSymbol(word) {
            word = word[:len(word)-1] // pop
        }
        fmt.Printf("%d: %s\n", i+1, word)
    }
}
