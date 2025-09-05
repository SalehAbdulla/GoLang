package main

import (
    "os"
    "fmt"
)

// $ go run . "(a)" "I'm heavy, jumpsuit is on steady, Lighter when I'm lower, higher when I'm heavy"
// 1: heavy
// 2: steady
// 3: heavy


func search(texts, exps  []string) (result []string) {
    for _, text := range texts {
        for _, exp := range exps {
            if contains(text, exp) {
                result = append(result, text)
            }
        }
    }
    return
}

func contains(text, toFind string) bool {
    for i := 0; i <= len(text) - len(toFind); i++ {
        if toFind == text[i:i+len(toFind)] {
            return true
        }
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
    
    if word != "" {
        words = append(words, word)
    }

    return words
}

func endsWithSymbol(s string) bool {
    if s == "" {return false}
    if s[len(s)-1] < 'a' || s[len(s)-1] > 'z' {return true}
    return false
}


func printAsNeeded(text []string) {
    for i, word := range text {
        if endsWithSymbol(word) {
            word = word[:len(word)-1]
        }
        fmt.Printf("%d: %s\n", i+1, word)
    }
}

func main(){

    args := os.Args[1:]
    if len(args) != 2 {return}
    regexp := args[0]
    
    if len(regexp) < 2 || regexp[0] != '(' || regexp[len(regexp)-1] != ')' {return}
    regexp = regexp[1:len(regexp)-1] // bracket removed

    var exps []string

    if contains(regexp, "|") {
        exps = append(exps, split(regexp, "|")...)
    } else {
        exps = append(exps, regexp)
    }
    
    // ---------------------------------------------

    text := args[1]

    var textToSlice []string
    textToSlice = append(textToSlice, split(text, " ")...)
    
    getResult := search(textToSlice, exps)

    printAsNeeded(getResult)

}
