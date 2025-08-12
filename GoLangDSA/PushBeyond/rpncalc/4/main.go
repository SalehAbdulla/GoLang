package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
)


func RpnCalc(s []string) (int, bool) {
    var stack []int
    for _, c := range s{
        cToN, err := strconv.Atoi(c)
        if err == nil {stack = append(stack, cToN); continue}
        if len(stack) < 2 {return -1, false}
        switch c {
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

func removeEmpty(s []string) (result []string) {
    for _, str := range s {
        if str != "" {
            result = append(result, str)
        }
    }
    return
}

func main(){
    args := os.Args[1:]
    if len(args) != 1 {fmt.Println("Error"); return}
    input := args[0]
    inputToSplit := strings.Split(input, " ")
    inputToSplit = removeEmpty(inputToSplit)

    getResult, ok := RpnCalc(inputToSplit)
    if !ok {fmt.Println("Error"); return}
    fmt.Println(getResult)
}
