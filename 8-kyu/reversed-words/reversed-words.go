package kata

import "strings"

func ReverseWords(str string) string {
	palavras := strings.Fields(str)
	for i, j := 0, len(palavras)-1; i < j; i, j = i+1, j-1 {
		palavras[i], palavras[j] = palavras[j], palavras[i]
	}
	return strings.Join(palavras, " ")
}
