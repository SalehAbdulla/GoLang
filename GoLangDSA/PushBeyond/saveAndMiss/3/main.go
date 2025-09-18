package piscine

func SaveAndMiss(arg string, num int) string {
    if num <= 0 {return arg}
    counter := 0
    flipper := true
    var result string

    for i := 0; i < len(arg); i++ {
        if counter == num {counter = 0; flipper = !flipper}
        if flipper {
            result += string(arg[i])
        }
        counter++
    }

    return result
}