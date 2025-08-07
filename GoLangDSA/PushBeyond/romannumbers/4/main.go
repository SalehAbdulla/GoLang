package main

import (
    "fmt"
    "os"
)

type RumanNumeral struct {
    Val int
    Symbol string
    Formula string
}

var numerals = []RumanNumeral {
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
    formula := ""
    symbol := ""
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

func main(){

    args := os.Args[1:]
    if len(args) != 1 {fmt.Println("ERROR: cannot convert to roman digit"); return}
    num, ok := StringToInt(args[0])
    if !ok {fmt.Println("ERROR: cannot convert to roman digit"); return}
    getFormula, getSymbol := GetRuman(num)
    fmt.Println(getFormula)
    fmt.Println(getSymbol)

}

func StringToInt(num string) (int, bool) {
    if num == "0" || num == "" || num[0] == '-' {return -1, false}

    result := 0
    for _, char := range num {
        if char >= 'a' && char <= 'z' {
            return -1, false
        } else {
            result *= 10
            result += int(char - '0')
        }
    }

    if result <= 0 || result > 3999 {return -1, false}
    return result, true
}


