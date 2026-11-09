package kata

import "fmt"

func SeriesSum(n int) string {
	sum := 0.0
	for i := range n {
		sum += 1.0 / float64(3*i+1)
	}
	return fmt.Sprintf("%.2f", sum)
}
