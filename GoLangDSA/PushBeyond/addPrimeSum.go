package main

import (
    "github.com/01-edu/z01"
    "os"
)

func main() {
    args := os.Args[1:]
    //If the number of arguments is different from 1, or if the argument is not a positive number, 
    //the program displays 0 followed by a newline.

    if len(args) != 1 {
        z01.PrintRune('0')
        z01.PrintRune('\n')
        return
    }

    number := atoi(args[0])
    if number <= 0 {
        z01.PrintRune('0')
        z01.PrintRune('\n')
        return
    }

    total := 0
    for i := 2; i <= number; i++ {
        if IsPrime(i) {
            total += i
        }
    }
    result := Itoa(total)
    for _, num := range result {
        z01.PrintRune(num)
    }
    z01.PrintRune('\n')

}


func IsPrime(num int) bool {
    if num < 2 {return false}
    for i := 2; i*i <= num; i++ {
        if num %i == 0 {
            return false
        }
    }
    return true
}

func atoi(number string) int {
    if number == "0" {return 0}
    isNegative := false
    if number[0] == '-' {
        isNegative = true
        number = number[1:]
    }

    result := 0

    for _, num := range number {
        if !(num >= '0' && num <= '9') {return 0}
        result *= 10
        result += int(num - '0')
    }

    if isNegative {
        return result * -1
    } else {
        return result
    }
}


func Itoa(number int) string {
    if number == 0 {return "0"}
    isNegative := false
    if number < 0 {
        isNegative = true
        number = -number
    }
    result := ""

    for number > 0 {
        digit := number % 10
        result = string(digit + '0') + result
        number = number / 10
    }


    if isNegative {
        return "-" + result 
    } else {
        return result
    }
}