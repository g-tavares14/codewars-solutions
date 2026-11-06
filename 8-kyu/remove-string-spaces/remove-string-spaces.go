package kata

func NoSpace(word string) string {
	var wordWithoutSpaces string

	for _, char := range word {
		if char != ' ' {
			wordWithoutSpaces += string(char)
		}
	}
	return wordWithoutSpaces
}
