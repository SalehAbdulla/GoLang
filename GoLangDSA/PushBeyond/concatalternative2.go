package piscine

func ConcatAlternate2(slice1,slice2 []int) []int {
    if len(slice2) > len(slice1) {
        slice1, slice2 = slice2, slice1
    }
    var result []int
    for i, char := range slice1 {
        result = append(result, char)
        if i < len(slice2) {
            result = append(result, slice2[i])
        }
    }   
    return result
}
