package kata

import "sort"

func TwoToOne(s1 string, s2 string) string {
	seen := make(map[rune]bool)

	for _, char := range s1 + s2 {
		seen[char] = true
	}

	var chars []rune
	for char := range seen {
		chars = append(chars, char)
	}

	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})

	return string(chars)

}
