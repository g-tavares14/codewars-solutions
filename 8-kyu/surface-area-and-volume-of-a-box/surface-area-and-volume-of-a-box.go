package kata

func GetSize(w, h, d int) [2]int {
	return [2]int{2 * (w*h + w*d + h*d), w * h * d}
}
