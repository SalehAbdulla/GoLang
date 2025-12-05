package main

import (
    "fmt"
    "os"
    "strconv"
    "strings"
)


func main(){
    args := os.Args[1:]
    if len(args) != 1 {fmt.Println("Error"); return}

    input := args[0]
    getCalc, ok := rpnCalc(input)
    if !ok {fmt.Println("Error"); return}
    fmt.Println(getCalc)
}


func rpnCalc(text string) (int, bool) {

    var stack []int
    textSplit := strings.Split(text, " ");

    for _, char := range textSplit {
        if char == "" {continue}
        
        charToNum, err := strconv.Atoi(char)
        if err == nil {stack = append(stack, charToNum); continue}

        if len(stack) < 2 {return -1, false}

        switch char {
            case "+":
                stack[len(stack)-2] += stack[len(stack)-1]
                stack = stack[:len(stack)-1] // pop
            case "-":
                stack[len(stack)-2] -= stack[len(stack)-1]
                stack = stack[:len(stack)-1] // pop
            case "/":
                stack[len(stack)-2] /= stack[len(stack)-1]
                stack = stack[:len(stack)-1] // pop
            case "*":
                stack[len(stack)-2] *= stack[len(stack)-1]
                stack = stack[:len(stack)-1] // pop
            case "%":
                stack[len(stack)-2] %= stack[len(stack)-1]
                stack = stack[:len(stack)-1] // pop
            default:
                return -1, false
        }

    }
    if len(stack) == 1 {
        return stack[0], true
    } 
    return -1, false
}



