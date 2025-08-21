package piscine

func SaveAndMiss(arg string, num int) string {
    if arg == "" || num == 0 {return arg}
    
    var result string
    saveMiss := true
    counter := 0
    for _, char := range arg {
        if saveMiss {result += string(char)}
        counter++
        if counter == num {saveMiss = !saveMiss; counter=0}
    }
    return result
}
