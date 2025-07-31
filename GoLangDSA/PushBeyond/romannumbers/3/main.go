package main

import (
    "os"
    "fmt"
)

type RumanType struct {
    Val int
    Symbol string
    Formula string
}

var numerals = []RumanType {
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

func getRuman(n int) (string, string) {
    symbol := ""
    formula := ""
    for _, numeral := range numerals {
        for n >= numeral.Val {
            symbol += numeral.Symbol
            if formula != "" {formula += "+"}
            formula += numeral.Formula
            n = n - numeral.Val
        }
    }
    return symbol, formula
}

func main(){
    args := os.Args[1:]
    if len(args) != 1 {fmt.Println("ERROR: cannot convert to roman digit"); return} 
    num, ok := StringToInt(args[0])
    if !ok {fmt.Println("ERROR: cannot convert to roman digit"); return} 
    getRumanSymbol, getRumanFormula := getRuman(num)
    fmt.Println(getRumanFormula)
    fmt.Println(getRumanSymbol)
}

func StringToInt(num string) (int, bool){
    result := 0
    for _, char := range num {
        if char < '0' || char > '9' {return -1, false}
        digit := int(char - '0')
        result *= 10
        result += digit
    }
    if result <= 0 || result > 3999 {return -1, false}
    return result, true
}