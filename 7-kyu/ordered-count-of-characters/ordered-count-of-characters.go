package orderedcount

func OrderedCount(text string) []Tuple {
	counts := make(map[rune]int)
	result := make([]Tuple, 0)

	for _, char := range text {
		counts[char]++
	}

	for _, char := range text {
		if count, ok := counts[char]; ok {
			result = append(result, Tuple{Char: char, Count: count})
			delete(counts, char)
		}
	}

	return result
}
