package piscine


func Slice(a []string, nbrs... int) []string{
    if len(a) == 0 || len(nbrs) == 0 {return a}
    starts := nbrs[0]
    if starts < 0 {starts += len(a)}

    ends := len(a)
    if len(nbrs) > 1 {
        ends = nbrs[1]
    }

    if ends < 0 {ends += len(a)}
    
    /// ------ validation

    if starts < 0 {starts = 0}
    if ends < 0 {ends = 0}
    if starts > len(a) {starts = len(a)}
    if ends > len(a) {ends = len(a)}
    if starts > ends {return nil}

    return a[starts:ends]




}
