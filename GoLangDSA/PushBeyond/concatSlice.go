package piscine

func ConcatSlice(slice1, slice2 []int) []int {
    if len(slice1) == 0 && len(slice2) == 0 {return nil}
    result := []int{}
    result = append(result, slice1 ...)
    result = append(result, slice2 ...)
    return result
}

