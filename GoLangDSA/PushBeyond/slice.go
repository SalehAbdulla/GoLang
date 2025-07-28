package piscine

func Slice(a []string, nbrs ...int) []string {

	l := len(a) // Get length
	if len(nbrs) == 0 {
		return nil
	} // Edge case

	// ----- get starts:ends
	starts := nbrs[0]
	if starts < 0 {
		starts += l
	}

	ends := l
	if len(nbrs) > 1 {
		ends = nbrs[1]
		if ends < 0 {
			ends += l
		}
	}

	//---- validation
	if starts < 0 {
		starts = 0
	}
	if ends < 0 {
		ends = 0
	}
	if starts > l {
		starts = l
	}
	if ends > l {
		ends = l
	}
	if starts > ends {
		return nil
	}

	// ----- return
	return a[starts:ends]
}
