package kata


func FindEvenIndex(arr []int) int {
	total := 0
	for _, v := range arr {
		total += v
	}

	sumLeft := 0
	for i, v := range arr {
		sumRight := total - sumLeft - v
		if sumLeft == sumRight {
			return i
		}
		sumLeft += v
	}

	return -1
}
