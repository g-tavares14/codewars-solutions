package kata

func Disemvowel(comment string) string {
	buf := make([]rune, 0, len(comment))

	for _, char := range comment {
		switch char {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			continue
		default:
			buf = append(buf, char)
		}
	}
	return string(buf)
}
