package kata

func FindOdd(seq []int) int {
	seen := make(map[int]int)

	for _, number := range seq {
		seen[number]++
	}

	for number, count := range seen {
		if count%2 == 1 {
			return number
		}
	}

	return 0
}
