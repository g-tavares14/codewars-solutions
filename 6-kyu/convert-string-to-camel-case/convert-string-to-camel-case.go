package kata

import (
	"regexp"
	"strings"
)

func ToCamelCase(s string) string {
	re := regexp.MustCompile(`[-_]`)
	parts := re.Split(s, -1)
	var result string

	for i, char := range parts {
		if i == 0 {
			result += char
		} else {
			result += strings.Title(char)
		}
	}

	return result

}
