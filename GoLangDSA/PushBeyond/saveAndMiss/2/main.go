package piscine


func SaveAndMiss(arg string, num int) string {
    if num < 1 {return arg}
    flipper := true
    counter := 0
    var buffer string
    for i := 0; i < len(arg); i++ {
        if counter == num {counter=0; flipper=!flipper}
        if flipper {
            buffer += string(arg[i])
        }
        counter++
    }
    return buffer
}
