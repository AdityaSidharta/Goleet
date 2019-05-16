package codewars

func NbYear(p0 int, percent float64, aug int, p int) int {
	count := 0
	percentage := 1. + percent/100.
	for p0 < p {
		p0 = int((float64(p0) * percentage) + float64(aug))
		count = count + 1
	}
	return count
}
