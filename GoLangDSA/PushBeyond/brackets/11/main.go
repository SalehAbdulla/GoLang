package main

import (
    "fmt"
    "os"
)

func BracketBalanced(str string) bool {
    var stack []rune
    for _, char := range str {
        if char == '[' || char == '(' || char == '{' {
            stack = append(stack, char)
        } else if char == ']' {
            if len(stack) == 0 || stack[len(stack)-1] != '[' {return false}
            stack = stack[:len(stack)-1]
        } else if char == '}' {
            if len(stack) == 0 || stack[len(stack)-1] != '{' {return false}
            stack = stack[:len(stack)-1]
        } else if char == ')' {
            if len(stack) == 0 || stack[len(stack)-1] != '(' {return false}
            stack = stack[:len(stack)-1]
        }
    }
    return len(stack) == 0
}   


func main(){
    args := os.Args[1:]
    if len(args) == 0 {fmt.Println(); return}
    for _, arg := range args {
        if BracketBalanced(arg) {
            fmt.Println("OK")
        } else {
            fmt.Println("Error")
        }
    }
}
