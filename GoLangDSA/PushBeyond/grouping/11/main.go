package main

import (
    "os"
    "fmt"
)

// $ go run . "(a)" "I'm heavy, jumpsuit is on steady, Lighter when I'm lower, higher when I'm heavy"

// 1: heavy
// 2: steady
// 3: heavy

func contains(s, toFind string) bool {
    for i := 0; i <= len(s) - len(toFind); i++ {
        if toFind == s[i:i+len(toFind)] {return true}
    }
    return false
}

func endsWithSymbol(s string) bool {
    l := s[len(s)-1]
    return !(l >= 'a' && l <= 'z') && !(l >= 'A' && l <= 'Z')
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

func PrintAsNeeded(result []string) {
    for i, str := range result {
        if endsWithSymbol(str) {
            str = str[:len(str)-1]
        }
        fmt.Printf("%d: %s\n", i+1, str)
    }
}

func main(){
    args := os.Args[1:]
    if len(args) != 2 {return}
    exp := args[0]
    
    if len(exp) < 3 || exp[0] != '(' || exp[len(exp)-1] != ')' {return}
    exp = exp[1:len(exp)-1]
    var expSlice []string

    if contains(exp, "|") {
        expSlice = append(expSlice, split(exp, "|")...)
    } else {
        expSlice = append(expSlice, exp)
    }
    
    text := args[1]
    var textSlice []string
    textSlice = append(textSlice, split(text, " ")...)

    getResult := search(expSlice, textSlice)
    PrintAsNeeded(getResult)

}