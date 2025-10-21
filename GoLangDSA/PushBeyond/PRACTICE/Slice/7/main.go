package piscine

func Slice(a []string, nbrs... int) []string{
    l := len(a)
    start := 0 
    if len(nbrs) > 0 {
        start = nbrs[0]
        if start < 0 {start += l}
    }

    end := l
    if len(nbrs) > 1 {
        end = nbrs[1]
        if end < 0 {end += l}
    }

    // ------ 5 Validation
    if start < 0 {start = 0}
    if end < 0 {end = 0}
    if start > l {start = l}
    if end > l {end = l}
    if start > end {return nil}

    return a[start:end]
}