package kata

import "strings"

func DuplicateEncode(word string) string {
	lower := strings.ToLower(word)
	count := make(map[rune]int)

	for _, char := range lower {
		count[char]++
	}

	var newString string

	for _, char := range lower {
		if count[char] == 1 {
			newString += "("
		} else {
			newString += ")"
		}
	}

	return newString

}
