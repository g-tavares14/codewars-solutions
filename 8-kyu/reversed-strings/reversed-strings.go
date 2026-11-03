package kata

func Solution(word string) string {
	left := 0
	right := len(word) - 1
	reversedWord := []rune(word)

	for left < right {
		reversedWord[left], reversedWord[right] = reversedWord[right], reversedWord[left]
		left++
		right--
	}
	return string(reversedWord)
}
