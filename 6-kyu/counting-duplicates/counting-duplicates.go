package kata

import "strings"

func duplicate_count(s1 string) int {
	stringLower := strings.ToLower(s1)
	count := make(map[rune]int)

	for _, char := range stringLower {
		count[char]++
	}

	duplicate := 0

	for _, value := range count {
		if value > 1 {
			duplicate++
		}
	}

	return duplicate
}
