package main

import (
    "fmt"
    "os"
)


func BracketMateched(s string) bool {
    var stack []rune

    sToRunes := []rune(s)

    for _, char := range sToRunes {
        if char == '{' || char == '(' || char == '[' {
            stack = append(stack, char)
        } else if char == '}' {
            if len(stack) == 0 || stack[len(stack) - 1] != '{' {return false}
            stack = stack[:len(stack)-1] // pop last Bracket
        } else if char == ')' {
            if len(stack) == 0 || stack[len(stack)- 1] != '(' {return false}
            stack = stack[:len(stack)-1] // pop last Bracket
        } else if char == ']' {
            if len(stack) == 0 || stack[len(stack)-1] != '[' {return false}
            stack = stack[:len(stack)-1] // pop last bracket 
        }
    }
    
    return len(stack) == 0
}



func main(){
    args := os.Args[1:]
    if len(args) == 0 {fmt.Println(); return}

    for _, arg := range args {
        if BracketMateched(arg) {
            fmt.Println("OK")
        } else {
            fmt.Println("Error")
        }
    }
}