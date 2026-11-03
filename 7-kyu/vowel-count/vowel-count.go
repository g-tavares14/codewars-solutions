package kata

import "strings"

func GetCount(str string) (count int) {
	var vowels = "aeiou"

	for _, char := range str {
		if strings.Contains(vowels, string(char)) {
			count++
		}
	}
	return count
}
