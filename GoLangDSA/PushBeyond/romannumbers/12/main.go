package main

import (
    "os"
    "fmt"
)


type RumanNumerals struct {
    Val int
    Symbol, Formula string
}

var numerals = []RumanNumerals {
    {1000,"M","M"},
    {900,"CM","(M-C)"},
    {500,"D","D"},
    {400,"CD","(D-C)"},
    {100,"C","C"},
    {90,"XC","(C-X)"},
    {50,"L","L"},
    {40,"XL","(L-X)"},
    {10,"X","X"},
    {9,"IX","(X-I)"},
    {5,"V","V"},
    {4,"IV","(V-I)"},
    {1,"I","I"},
}

func GetRuman(num int) (string, string) {
    var formula string
    var symbol string
    for _, numeral := range numerals {
        for num >= numeral.Val {
            symbol += numeral.Symbol
            if formula != "" {formula += "+"}
            formula += numeral.Formula
            num -= numeral.Val
        }
    }
    return formula, symbol
}

func StringToInt(num string) (int, bool) {
    if len(num) == 0 || num == "0" {return -1, false} 
    if num[0] == '-' {return -1, false}
    var result int
    for _, char := range num {
        if char < '0' || char > '9' {
            {return -1, false} 
        } else {
            result *= 10
            result += int(char - '0')
        }
    }

    if result < 1 || result > 3999 {return -1, false}

    return result, true

}


func main(){
    args := os.Args[1:]
    if len(args) != 1 {return}
    input, pass := StringToInt(args[0])
    if !pass {fmt.Println("ERROR: cannot convert to roman digit"); return}
    getFormula, getSymbol := GetRuman(input)
    fmt.Println(getFormula)
    fmt.Println(getSymbol)
}




