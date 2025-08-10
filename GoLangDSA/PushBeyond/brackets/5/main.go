package main

import (
    "os"
    "fmt"
)

// ---------------------------

func BracketBalancedIn(s string) bool {
    var stack []rune
    for _, char := range s {
        if char == '[' || char == '{' || char == '(' {
            stack = append(stack, char)
        } else if char == ']' {
            if len(stack) == 0 || stack[len(stack)-1] != '[' {return false}
            stack = stack[:len(stack)-1] // pop
        } else if char == ')' {
            if len(stack) == 0 || stack[len(stack)-1] != '(' {return false}
            stack = stack[:len(stack)-1] // pop
        } else if char == '}' {
            if len(stack) == 0 || stack[len(stack)-1] != '{' {return false}
            stack = stack[:len(stack)-1] // pop
        }
    }
    return len(stack) == 0
}

// ---------------------------

func main(){
    args := os.Args[1:]
    if len(args) == 0 {fmt.Println(); return}
    for _, arg := range args{
        if BracketBalancedIn(arg) {
            fmt.Println("OK")
        } else {
            fmt.Println("Error")
        }
    }
}