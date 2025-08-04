package main

import (
    "os"
    "fmt"
)

// Handle Brackets - remove them, return []string
func getExpSlice(exp string) (result []string) {
    // (s|a) -- if contains | -- then return {"s", "a"} 
    if contains(exp, "|") {
        result = append(result, split(exp[1:len(exp)-1], "|")...) // remove the Bracket, & split
    } else {
        // else (a) -- then return {"a"}
        result = append(result, string(exp[1:len(exp)-1])) // remove the Bracket, and return
    }
    return
}

func split(s, splitter string) (words []string) {
    word := ""
    for _, char := range s {
        if string(char) != splitter {
            word += string(char)
        } else {
            words = append(words, word)
            word = ""
        }
    }
    if word != "" {words = append(words, word)}
    return 
}

func contains(s, toFind string) bool {
    for i := 0; i <= len(s)-len(toFind); i++ {
        if toFind == s[i:i+len(toFind)] {return true}
    }
    return false
}

func searchInText(allExp []string, texts string) (result []string) {
    for _, text := range split(texts, " ") {
        for _, exp := range allExp {
            if contains(text, exp) {
                result = append(result, text)
            }
        }
    }
    return
}

func endsWithRune(str string) bool {
    lastRune := str[len(str)-1]
    return (lastRune >= 'a' && lastRune <= 'z') || (lastRune >= 'A' && lastRune <= 'Z')
}

func main(){

    args := os.Args[1:]
    if len(args) == 2 {

        regularExp := args[0]
        textToSearch := args[1]
        if regularExp == "" || textToSearch == "" {return}
        if regularExp[0] != '(' && regularExp[len(regularExp)-1] != ')' {return}

        searchResult := searchInText(getExpSlice(regularExp), textToSearch)
        for i, str := range searchResult {
            if !endsWithRune(str) {
                fmt.Printf("%d: %s\n", i+1, str[:len(str)-1]) // pop last rune, probably ,
            } else {
                fmt.Printf("%d: %s\n", i+1, str)
            }
        }
    }

}