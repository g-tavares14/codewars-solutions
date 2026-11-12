package kata

func TwoSum(numbers []int, target int) [2]int {
	seen := make(map[int]int)

	for i, num := range numbers {
		complemento := target - num
		if j, ok := seen[complemento]; ok {
			return [2]int{j, i}
		}
		seen[num] = i
	}
	return [2]int{}
}
