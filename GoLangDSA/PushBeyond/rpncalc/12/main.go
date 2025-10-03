package main

import (
    "strconv"
    "os"
    "fmt"
    "strings"
)

func CalcRpn(strs []string) (int, bool) {
    var stack []int
    for _, str := range strs {
        if str == "" {continue}
        strInt, err := strconv.Atoi(str)
        if err == nil {stack = append(stack, strInt); continue}
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
    }
    return -1, false
}

func main(){
    args := os.Args[1:]
    if len(args) != 1 {fmt.Println("Error"); return}
    if len(args[0]) == 1 {fmt.Println(args[0]); return}
    inputSplit := strings.Split(args[0], " ")


    getResult, ok := CalcRpn(inputSplit)
    if !ok {fmt.Println("Error"); return}

    fmt.Println(getResult)

}