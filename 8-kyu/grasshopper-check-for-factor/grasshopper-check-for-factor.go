package kata

func CheckForFactor(base int, factor int) bool {
	if base%factor == 0 {
		return true
	}
	return false
}
