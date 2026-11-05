package kata

func SimpleMultiplication(n int) int {
	var sum int

	if n%2 == 0 {
		sum = n * 8
	} else {
		sum = n * 9
	}

	return sum

}
