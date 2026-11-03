package kata

func SmallestIntegerFinder(numbers []int) int {
	smallValue := numbers[0]

	for i := 0; i < len(numbers); i++ {
		if numbers[i] < smallValue {
			smallValue = numbers[i]
		}
	}
	return smallValue
}
