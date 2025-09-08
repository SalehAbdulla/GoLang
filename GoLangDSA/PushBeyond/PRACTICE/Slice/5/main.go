package picsine


func Slice(a []string, nbrs... int) []string {
    if len(a) == 0 || len(nbrs) == 0 {return a}

    start := nbrs[0]
    if start < 0 {start += len(a)}

    end := len(a)
    if len(nbrs) > 1 {end = nbrs[1]} 
    if end < 0 {end += len(a)}

    // validation
    if start < 0 {start = 0}
    if end < 0 {end = 0}
    if start > len(a) {start = len(a)}
    if end > len(a) {end = len(a)}
    if start > end {return nil}

    return a[start:end]
}
