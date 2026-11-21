package kata

import "fmt"

func HoopCount(n int) string {
	if n < 10 {
		return fmt.Sprintf("Keep at it until you get it")
	}
	return fmt.Sprintf("Great, now move on to tricks")
}
