package piscine

func Slice(a []string, nbrs... int) []string{
    if len(nbrs) == 0 {return a}
    l := len(a)
    starts := nbrs[0]
    
    if starts < 0 {starts += l}
    ends := l

    if len(nbrs) > 1 {
        ends = nbrs[1]
        if ends < 0  {ends += l}
    }

    /// ---- 5 Validation

    if starts > l {starts = l}
    if ends > l {ends = l}
    if starts < 0 {starts = 0}
    if ends < 0 {ends = 0}
    if starts > ends {return nil}

    // --------------------

    return a[starts:ends]


}
