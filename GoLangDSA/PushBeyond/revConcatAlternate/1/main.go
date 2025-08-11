package piscine

func RevConcatAlternate(slice1,slice2 []int) []int {
    if len(slice1) == 0 && len(slice2) == 0 {return []int{}}
    var result []int
    r1, r2 := reverse(slice1), reverse(slice2)
    if len(r1) > len(r2) {
        result = append(result, r1[:len(r1)-len(r2)]...)
        r1 = r1[len(r1)-len(r2):]
    } else if len(r2) > len(r1) {
        result = append(result, r2[:len(r2)-len(r1)]...)
        r2 = r2[len(r2)-len(r1):]
        
    }
    for i := 0; i < len(r1); i++ {
        result = append(result, r1[i], r2[i])
    }
    return result
}


func reverse(s []int) (result []int) {
    for i := range s {
        result = append(result, s[len(s)-1-i])
    }
    return
}