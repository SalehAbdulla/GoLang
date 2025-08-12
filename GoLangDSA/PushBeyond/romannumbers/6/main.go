package main

import("os"; "fmt")

type RumanNumeral struct {
    val int
    symbol string
    formula string
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

func GetRuman(n int) (string, string) {
    var symbol string
    var formula string
    for _, numeral := range numerals {
        for n >= numeral.val {
            symbol += numeral.symbol
            if formula != "" {formula += "+"}
            formula += numeral.formula
            n -= numeral.val
        }
    }
    return symbol, formula
}

func main(){
    args := os.Args[1:]
    if len(args) != 1 {return}
    input := args[0]

    inputToInt, ok := StringToInt(input)
    if !ok {fmt.Printf("%s\n", "ERROR: cannot convert to roman digit");return}

    getSymbol, getFormula := GetRuman(inputToInt)
    fmt.Println(getFormula)
    fmt.Println(getSymbol)



}

func StringToInt(num string) (int, bool) {
    if num == "" || num == "0" || num[0] == '-' {return -1, false}
    result := 0
    for _, char := range num {
        result *= 10
        result = int(char - '0') + result
    }
    if result < 2 || result > 3999 {return -1, false}
    return result, true
}








