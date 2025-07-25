package piscine

func SaveAndMiss(arg string, num int) string {
    ///If the int is 0 or a negative number return the original string.
    if num == 0 {return arg}

    result := ""
    save := true

    for i := 0; i < len(arg); i += num {
        end := i + num // get slice value
        if end > len(arg) {end = len(arg)} // In case outOfRange
        if save {result += arg[i:end]} // only on save mode
        save = !save // Flip
    }
    
    return result
}