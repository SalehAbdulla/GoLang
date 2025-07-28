package piscine

func RevConcatAlternate(slice1,slice2 []int) []int {
    // 1) Reverse slices
    // 2) r1[:len(r1)-len(r2)...] first elements
    // 3) get rid of r1 = r1[len(r1)-len(r2):]
    // 4) add them alternativly append(res, r1[i], r2[i])
    
    r1, r2 := reverse(slice1), reverse(slice2)
    res := []int{}

    if len(r1) > len(r2) {
        res = append(res, r1[:len(r1)-len(r2)]...)
        r1 = r1[len(r1)-len(r2):]
    } else if len(r2) > len(r1) {
        res = append(res, r2[:len(r2)-len(r1)]...)
        r2 = r2[len(r2)-len(r1):]
    }

    for i := 0; i < len(r1); i++ {
        res = append(res, r1[i], r2[i])
    }

    return res
}

func reverse(slice []int) []int{
    res := []int{}
    for i := range slice {
        res = append(res, slice[len(slice)-1-i])
    }
    return res
}