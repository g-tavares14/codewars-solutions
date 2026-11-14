package kata

func ToAlternatingCase(str string) string {
	var result string

	for _, char := range str {
		if char >= 'a' && char <= 'z' {
			result += string(char - 32)
		} else if char >= 'A' && char <= 'Z' {
			result += string(char + 32)
		} else {
			result += string(char)
		}
	}

	return result
}
