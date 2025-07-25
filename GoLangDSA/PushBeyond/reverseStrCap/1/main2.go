package main

import (
    "os"
    "github.com/01-edu/z01"
)

func main(){
    // If there are no argument, the program displays nothing.
    args := os.Args[1:]
    if len(args) == 0 {return}
    if args[0] == "" {
        z01.PrintRune('\n')
        return
    } 
    
    // 1) Lower case everything, save the result []string{}
    everythingSmall := []string{}
    temp := ""
    for _, words := range args {
        for _, char := range words {
            if char >= 'A' && char <= 'Z' {
                temp += string(rune(char + 32)) // lower case
                continue
            }
            temp += string(char)
        }
        // after word finishes
        if temp != "" {
            everythingSmall = append(everythingSmall, temp)
            temp = ""
        }
    }

    // 2) Loop through each string, char before a space make it upper, save it
    result := []string{}
    temp2 := ""
    for _, word := range everythingSmall {
        for i, char := range word {
           if i + 1 < len(word) && rune(word[i + 1]) == ' ' || i == len(word) - 1{
                if char >= 'a' && char <= 'z' {
                    temp2 += string(rune(char - 32)) // upper case
                    continue
                }
           } 
           temp2 += string(char)
        }
        if temp2 != "" {
            result = append(result, temp2)
            temp2 = ""
        }
    }

    // 3) print the []string, with \n in between
    for _, word := range result {
        for _, char := range word {
            z01.PrintRune(char)
        }
        z01.PrintRune('\n')
    }
}