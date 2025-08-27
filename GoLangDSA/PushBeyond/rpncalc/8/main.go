package main

import (
    "os"
    "fmt"
    "strings"
    "strconv"
)

func RpnCalc(strs []string) (int, bool) {
    var stack []int
    for _, str := range strs {
        if str == "" || str == " " {continue}
        strToInt, err := strconv.Atoi(str)
        if err == nil {stack = append(stack, strToInt); continue}
        if len(stack) < 2 {return -1, false}
        switch str {
            case "+":
                stack[len(stack)-2] += stack[len(stack)-1]
                stack = stack[:len(stack)-1] // pop
            case "-":
                stack[len(stack)-2] -= stack[len(stack)-1]
                stack = stack[:len(stack)-1] // pop
            case "*":
                stack[len(stack)-2] *= stack[len(stack)-1]
                stack = stack[:len(stack)-1] // pop
            case "/":
                stack[len(stack)-2] /= stack[len(stack)-1]
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
    } else {
        return -1, false
    }
}


func main(){
    args := os.Args[1:]
    if len(args) != 1 {fmt.Println("Error"); return}
    // Split, then remove empty
    splitInput := strings.Split(args[0], " ")
    getResult, ok := RpnCalc(splitInput)
    if !ok {fmt.Println("Error"); return}
    fmt.Println(getResult)
}
