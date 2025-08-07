package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
)

func clearEmpty(s []string) (result []string) {
    for _, char := range s {
        if char != "" {
            result = append(result, string(char))
        }
    }
    return
}

func main(){
    args := os.Args[1:]
    if len(args) != 1 {fmt.Println("Error"); return}
    input := strings.Split(args[0], " ")
    input = clearEmpty(input)

    var stack []int
    for _, str := range input {

        num, err := strconv.Atoi(str)
        if err == nil {stack = append(stack, num); continue}

        // We are expecting an operator, and two digits in stack
        if len(stack) < 2 {fmt.Println("Error"); return}
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
                fmt.Println("Error")
                return
        }
    }

    if len(stack) == 1 {
        fmt.Println(stack[0])
    } else {
        fmt.Println("Error")
    }


}   