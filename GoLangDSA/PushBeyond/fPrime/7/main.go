package main

import ("os"; "fmt"; "strconv")
// $ go run . 225225
// 3*3*5*5*7*11*13

func main() {
    args := os.Args[1:]
    if len(args) != 1 {return}

    num, err := strconv.Atoi(args[0])
    if err != nil || num < 2 {return}


    var result []int
    i := 2
    for num >= i {
        for num %i == 0{
            result = append(result, i)
            num /= i
        }
        i++
    }
    
    for j, n := range result {
        fmt.Print(n) 
        if j != len(result) - 1 {
            fmt.Print("*")
        }
    }
    fmt.Println()
}