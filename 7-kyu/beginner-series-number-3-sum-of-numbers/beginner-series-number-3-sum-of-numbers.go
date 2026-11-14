package kata


func GetSum(a, b int) int {
	lo, hi := a, b
	if lo > hi {
		lo, hi = hi, lo
	}
	return (lo + hi) * (hi - lo + 1) / 2
}
