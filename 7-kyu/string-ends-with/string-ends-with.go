package kata

func solution(str, ending string) bool {
	if len(ending) > len(str) {
		return false
	}
	offset := len(str) - len(ending)
	for i := 0; i < len(ending); i++ {
		if str[offset+i] != ending[i] {
			return false
		}
	}
	return true
}
