package leetcode

func getMin(values []int) int {
	if len(values) == 0 {
		return 0
	} else {
		final := values[0]
		for idx := 1; idx < len(values); idx++ {
			cur := values[idx]
			if cur < final {
				final = cur
			}
		}
		return final
	}
}

func allSame(values []string, bigidx int) bool {
	if len(values) == 0 {
		return false
	} else {
		value := values[0][bigidx]
		for idx := 1; idx < len(values); idx++ {
			cur := values[idx][bigidx]
			if value != cur {
				return false
			}
		}
		return true
	}
}

func longestCommonPrefix(strs []string) string {
	nstrs := len(strs)
	if nstrs == 0 {
		return ""
	}
	nstrsvalue := make([]int, nstrs)
	for idx := 0; idx < nstrs; idx++ {
		nstrsvalue[idx] = len(strs[idx])
	}
	minstrsvalue := getMin(nstrsvalue)
	result := 0
	for idx := 0; idx < minstrsvalue; idx++ {
		if allSame(strs, idx) {
			result = result + 1
		} else {
			break
		}
	}
	return string(strs[0][:result])
}
