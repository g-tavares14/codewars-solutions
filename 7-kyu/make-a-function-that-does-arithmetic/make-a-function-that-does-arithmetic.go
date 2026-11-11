package kata

func Arithmetic(a int, b int, operator string) int {
	switch operator {
	case "add":
		return a + b
	case "subtract":
		return a - b
	case "multiply":
		return a * b
	case "divide":
		if b != 0 {
			return a / b
		}
		return 0
	default:
		return 0
	}
}
