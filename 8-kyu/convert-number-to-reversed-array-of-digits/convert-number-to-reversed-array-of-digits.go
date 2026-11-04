package kata

import "fmt"

func Digitize(n int) []int {
	var digits []int
	for _, digit := range fmt.Sprint(n) {
		digits = append(digits, int(digit-'0'))
	}

	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}
	return digits
}
