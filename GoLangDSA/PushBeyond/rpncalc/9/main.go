package main

import (
    "fmt"
    "os"
    "strconv"
    "strings"
)

//$ go run . "1 2 * 3 * 4 +" | cat -e
// 10$
func removeEmpty(text []string) (result []string) {
    for _, txt := range text {
        if txt == ""  || txt == " " {
            continue
        } else {
            result = append(result, txt)
        }
    }
    return
}

func RpnCalc(text []string) (int, bool) {
    var stack []int
    for _, str := range text {
        if str == "" || str == " " {continue}
        strToInt, err := strconv.Atoi(str)
        if err == nil {stack = append(stack, strToInt); continue}
        if len(stack) < 2 {return -1, false}
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
    } else {
        return -1, false
    }
}


func main() {

    args := os.Args[1:]
    if len(args) != 1 {fmt.Println("Error"); return}
    if len(args[0]) == 1 {fmt.Println(args[0]); return}
    text := strings.Split(args[0], " ")
    getResult, ok := RpnCalc(text)
    if !ok {fmt.Println("Error"); return}

    fmt.Println(getResult)

}