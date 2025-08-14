package main

import (
    "fmt"
    "os"
)

func Contains(text, toFind string) bool {
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

func search(texts, exps []string) (result []string) {
    for _, text := range texts {
        for _, exp := range exps {
            if Contains(text, exp) {
                result = append(result, text)
            }
        }
    }
    return
}

func endsWithSymbol(s string) bool {
    return !(s[len(s)-1] >= 'a' && s[len(s)-1] <= 'z') && !(s[len(s)-1] >= 'A' && s[len(s)-1] <= 'Z')
}

func main(){

    args := os.Args[1:]
    if len(args) != 2 {return}
    regexp := args[0]

    if len(regexp) < 2 || regexp == "" {return}
    if regexp[0] != '(' && regexp[len(regexp)-1] != ')' {return}

    regexp = regexp[1:len(regexp)-1] // removing ()
    text := args[1]
    if text == "" {return}

    var expToSlice []string

    if Contains(regexp, "|") {
        expToSlice = append(expToSlice, split(regexp, "|")...)
    } else {
        expToSlice = append(expToSlice, regexp)
    }

    // --------
    textToSlice := split(text, " ")

    getResult := search(textToSlice, expToSlice)

    for i, word := range getResult {
        if endsWithSymbol(word) {
            word = word[:len(word)-1]
        }

        fmt.Printf("%d: %s\n", i+1, word)

    }

}