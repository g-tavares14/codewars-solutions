package kata

func CountPositivesSumNegatives(numbers []int) []int {
	var res []int

	if len(numbers) == 0 {
		return res
	}

	var countPositives int
	var sumNegatives int

	for _, num := range numbers {
		if num > 0 {
			countPositives++
		} else if num < 0 {
			sumNegatives += num
		}

	}
	return []int{countPositives, sumNegatives}
}
