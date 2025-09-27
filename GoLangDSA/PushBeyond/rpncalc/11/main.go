// The Price is easy if the promise is a clear
package main

import (
    "fmt"
    "strconv"
    "os"
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
                stack[len(stack) - 2] += stack[len(stack) - 1]
                stack = stack[:len(stack)-1] // pop
            case "-":
                stack[len(stack)-2] -= stack[len(stack)-1]
                stack = stack[:len(stack)-1]
            case "*":
                stack[len(stack)-2] *= stack[len(stack)-1]
                stack = stack[:len(stack)-1]
            case "/":
                stack[len(stack)-2] /= stack[len(stack)-1]
                stack = stack[:len(stack)-1]
            case "%":
                stack[len(stack)-2] %= stack[len(stack)-1]
                stack = stack[:len(stack)-1]
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
    if len(args) != 1 || args[0] == "" {fmt.Println("Error"); return}
    inputSlice := strings.Split(args[0], " ")
    result, ok := CalcRpn(inputSlice)
    if !ok {fmt.Println("Error"); return}
    fmt.Println(result)
}

