package piscine

func Slice(a []string, nbrs ...int) []string {
	if len(a) == 0 || len(nbrs) == 0 {
		return a
	}
	l := len(a)

	start := nbrs[0]
	if start < 0 {
		start += l
	}

	end := len(a)
	if len(nbrs) > 1 {
		end = nbrs[1]
	}
	if end < 0 {
		end += l
	}

	// ---- validtion
	if start < 0 {
		start = 0
	}
	if end < 0 {
		end = 0
	}
	if start > l {
		start = l
	}
	if end > l {
		end = l
	}
	if start > end {
		return nil
	}

	return a[start:end]
}
