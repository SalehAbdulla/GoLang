package piscine


// The new slice should start with an element of the largest slice.

func ConcatAlternate(slice1, slice2 []int) []int {

    if len(slice1) == 0 && len(slice2) == 0 {return nil}

    if len(slice1) == len(slice2) {
        result := []int{}
        for _, num := range slice1 {result = append(result, num)}
        for _, num := range slice2 {result = append(result, num)}
    }

    s1Length := len(slice1)
    s2Length := len(slice2)

    // Always starts with slice1.
    // Here I'm shifting the slice1 in case its smaller
    if s1Length < s2Length {
        slice1, slice2 = slice2, slice1
    }

    result := []int{}
    // Make a for loop 
    for i := 0; i < len(slice1); i++ {
        if i < len(slice1) {result = append(result, slice1[i])}
        if i < len(slice2) {result = append(result, slice2[i])}
    }

    return result

}
