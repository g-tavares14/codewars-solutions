package kata

func Points(games []string) int {
	points := 0
	for _, game := range games {
		x, y := game[0], game[2]
		switch {
		case x > y:
			points += 3
		case x == y:
			points++
		}
	}
	return points
}
