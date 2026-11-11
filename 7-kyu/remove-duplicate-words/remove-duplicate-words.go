package kata

import "strings"

func RemoveDuplicateWords(str string) string {
	wordMap := make(map[string]bool)
	var result []string

	for word := range strings.FieldsSeq(str) {
		if !wordMap[word] {
			result = append(result, word)
			wordMap[word] = true
		}
	}

	return strings.Join(result, " ")
}
