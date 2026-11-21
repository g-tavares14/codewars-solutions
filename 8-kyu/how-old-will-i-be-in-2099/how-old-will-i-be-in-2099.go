package kata

import "fmt"

func pluralize(n int) string {
	if n == 1 {
		return "year"
	}
	return "years"
}

func calculateAge(born, year int) string {
	diff := year - born

	switch {
	case diff > 0:
		return fmt.Sprintf("You are %d %s old.", diff, pluralize(diff))
	case diff < 0:
		diff = -diff
		return fmt.Sprintf("You will be born in %d %s.", diff, pluralize(diff))
	default:
		return "You were born this very year!"
	}

}
