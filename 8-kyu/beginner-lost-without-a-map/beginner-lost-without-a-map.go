package kata

func Maps(x []int) []int {
	var result []int

	for _, num := range x {
		result = append(result, num*2)
	}
	return result

}
