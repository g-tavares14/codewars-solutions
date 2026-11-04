package kata

func FindOutlier(integers []int) int {
	totalEven := 0
	totalOdd := 0

	for _, value := range integers {
		if value%2 == 0 {
			totalEven++
		} else {
			totalOdd++
		}
	}

	switch {
	case totalEven > totalOdd:
		for _, value := range integers {
			if value%2 != 0 {
				return value
			}
		}
	case totalOdd > totalEven:
		for _, value := range integers {
			if value%2 == 0 {
				return value
			}
		}
	}

	return 0
}
