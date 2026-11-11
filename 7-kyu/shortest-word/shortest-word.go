package kata

import "strings"

func FindShort(s string) int {
	words := strings.Fields(s)
	if len(words) == 0 {
		return 0
	}

	shortest := len(words[0])
	for _, word := range words[1:] {
		if len(word) < shortest {
			shortest = len(word)
		}
	}
	return shortest
}
