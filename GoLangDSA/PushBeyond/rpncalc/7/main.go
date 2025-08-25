package main

import ("os";"fmt";"strconv";"strings")

func RpnCalc(strs []string) (int, bool) {
    var stack []int
    for _, str := range strs {
        str = strings.TrimSpace(str)
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

func main() {
    args := os.Args[1:]
    if len(args) != 1 {fmt.Println("Error"); return}
    inputToStr := strings.Split(args[0], " ")
    calc, ok := RpnCalc(inputToStr)
    if !ok {fmt.Println("Error"); return}
    fmt.Println(calc)
}
