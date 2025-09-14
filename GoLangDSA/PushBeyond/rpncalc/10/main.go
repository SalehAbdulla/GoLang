package main

import (
    "os"
    "fmt"
    "strconv"
    "strings"
)

func RpnCalc(strs []string) (int, bool) {
    var stack []int
    for _, str := range strs {
        if str == "" || str == " " {continue}
        strToInt, err := strconv.Atoi(str)
        if err == nil {stack = append(stack, strToInt); continue}
        // if not a number, we expecting two digits in the stack,
        // and a valid operator +-*%/
        if len(stack) < 2 {return -1, false}
        switch str {
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


func main() {
    args := os.Args[1:]
    if len(args) != 1 {fmt.Println("Error"); return}
    input := args[0]
    inputSlice := strings.Split(input, " ")
    calc, pass := RpnCalc(inputSlice)
    if !pass {fmt.Println("Error"); return}
    fmt.Println(calc)
}