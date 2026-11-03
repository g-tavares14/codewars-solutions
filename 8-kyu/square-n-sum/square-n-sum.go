package kata

func SquareSum(numbers []int) int {
  sum := 0

	for _, value := range numbers {
		sum += value * value

	}
	
	return sum
    
}
