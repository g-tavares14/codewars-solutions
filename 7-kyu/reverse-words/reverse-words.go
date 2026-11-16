package kata

import "strings"

func ReverseWords(str string) string {
	words := strings.Split(str, " ")
	for k, w := range words {
		r := []rune(w)
		for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
			r[i], r[j] = r[j], r[i]
		}
		words[k] = string(r)
	}
	return strings.Join(words, " ")
}
