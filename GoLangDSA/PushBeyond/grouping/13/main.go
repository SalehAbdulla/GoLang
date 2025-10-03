package main

import (
    "os"
    "fmt"
)


func contains(text, toFind string) bool {
    for i := 0; i <= len(text) - len(toFind); i++ {
        if toFind == text[i:i+len(toFind)]{return true}
    }
    return false
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

func printAsNeeded(result []string) {
    for i, str := range result {
        if !endsWithChar(str) {str = str[:len(str)-1]} // pop
        fmt.Printf("%d: %s\n", i+1, str)
    }
}

func endsWithChar(str string) bool {
    l := len(str) - 1
    return str[l] >= 'a' && str[l] <= 'z' || str[l] >= 'A' && str[l] <= 'Z'
}

func split(text, splitter string) (result []string) {
    var buffer string
    for _, char := range text {
        if string(char) != splitter {
            buffer += string(char)
        } else if buffer != "" {
            result = append(result, buffer)
            buffer = ""
        }
    }
    if buffer != "" {result = append(result, buffer)}
    return
}

// if the regular expression is not valid,
func main(){
    args := os.Args[1:]
    if len(args) != 2 {return}
    regexp := args[0]
    if len(regexp) < 3 {return}
    text := args[1]
    if text == "" {return}
    // --------------------- Edge cases ^
    regexp = regexp[1:len(regexp)-1]
    var regexpSlice []string
    if contains(regexp, "|") {
        regexpSlice = append(regexpSlice, split(regexp, "|")...)
    } else {
        regexpSlice = append(regexpSlice, regexp)
    }
    // --------------
    textSlice := split(text, " ")
    // --------------
    getSearchResult := search(regexpSlice, textSlice)
    printAsNeeded(getSearchResult)

}
