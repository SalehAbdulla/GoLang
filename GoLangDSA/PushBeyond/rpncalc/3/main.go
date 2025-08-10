package main

import (
    "fmt"
    "os"
    "strconv"
    "strings"
)

func removeSpaces(s []string) (result []string) {
    for _, str := range s {
        if str == " " || str == "" {
            continue
        } else {
            result = append(result, str)
        }
    }
    return
}

func CalcRpn(input string) (result int) {
    input2 := strings.Split(input, " ")
    input2 = removeSpaces(input2)
    var stack []int
    for _, str := range input2 {
        strToInt, err := strconv.Atoi(str)
        if err == nil {stack = append(stack, strToInt); continue}
        // If we didn't encountered a number, then we expect 
        // an operator and two digits in the stack
        if len(stack) < 2 {return -1}
        switch str {
            case "+":
                stack[len(stack)-2] += stack[len(stack)-1]
                stack = stack[:len(stack)-1]
            case "-":
                stack[len(stack)-2] -= stack[len(stack)-1]
                stack = stack[:len(stack)-1]
            case "*":
                stack[len(stack)-2] *= stack[len(stack)-1]
                stack = stack[:len(stack)-1]
            case "%":
                stack[len(stack)-2] %= stack[len(stack)-1]
                stack = stack[:len(stack)-1]
            case "/":
                stack[len(stack)-2] /= stack[len(stack)-1]
                stack = stack[:len(stack)-1]
            default:
                return -1
        }
    }

    if len(stack) == 1 {
        return stack[0]
    } else {
        return -1
    }

}


func main(){
    arg := os.Args[1:]
    if len(arg) != 1 || arg[0] == "" {fmt.Println("Error"); return}
    result := CalcRpn(arg[0])
    if result == -1 { // This logic should be change, maybe we got an operation with equals to -1
        fmt.Println("Error")
    } else {
        fmt.Println(result)
    }
}