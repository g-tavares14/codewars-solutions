package kata

func NumberOfPairs(gloves []string) int {
	listPairs := make(map[string]int)

	for _, glove := range gloves {
		color := glove
		listPairs[color]++
	}

	var countPairs int

	for _, pair := range listPairs {
		countPairs += pair / 2
	}

	return countPairs
}
