package piscine



func Slice(a []string, nbrs... int) []string{
    if len(nbrs) == 0 {return a}
    starts := nbrs[0]
    if starts < 0 {starts += len(a)}
    
    l := len(a)
    
    ends := l
    if len(nbrs) > 1 {ends = nbrs[1]}
    if ends < 0 {ends += len(a)}

    // --- Validation
    if starts < 0 {starts = 0}
    if ends < 0 {ends = 0}
    if starts > l {starts = l}
    if ends > l {ends = l}
    if starts > ends {return nil}

    return a[starts:ends]
}