package main

import (
    "os"
    "fmt"
)

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
    if word != "" {result = append(result, word)}
    return
}

func Contains(text, toFind string) bool {
    for i := 0; i <= len(text) - len(toFind); i++ {
        if text[i:i+len(toFind)] == toFind {
            return true
        }
    }
    return false
}

func EndsWithSymbol(t string) bool {
    return (t[len(t)-1] >= 'a' && t[len(t)-1] <= 'z') || (t[len(t)-1] >= 'A' && t[len(t)-1] <= 'Z')
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

func main(){
    args := os.Args[1:]
    if len(args) != 2 {return}
    exp := args[0]
    text := args[1]

    if text == "" {return}
    if exp[0] != '(' || exp[len(exp)-1] != ')' {return}
    exp = exp[1:len(exp)-1] // remove ()

    var expSlice []string
    if Contains(exp, "|") {
        expSlice = append(expSlice, Split(exp, "|")...)
    } else {
        expSlice = append(expSlice, exp)
    } 

    // -------------
    solution := Search(expSlice, Split(text, " "))
    for i, w := range solution {
        if !EndsWithSymbol(w) {
            w = w[:len(w)-1] // pop
        }
        fmt.Printf("%d: %s\n", i+1, w)
    }



}