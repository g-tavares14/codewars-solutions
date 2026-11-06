package kata

func Snail(snailMap [][]int) []int {
	if len(snailMap) == 0 || len(snailMap[0]) == 0 {
		return []int{}
	}

	top, bottom := 0, len(snailMap)-1
	left, right := 0, len(snailMap[0])-1
	res := make([]int, 0, len(snailMap)*len(snailMap[0]))

	for top <= bottom && left <= right {

		for j := left; j <= right; j++ {
			res = append(res, snailMap[top][j])
		}
		top++

		for i := top; i <= bottom; i++ {
			res = append(res, snailMap[i][right])
		}
		right--

		if top <= bottom {
			for j := right; j >= left; j-- {
				res = append(res, snailMap[bottom][j])
			}
			bottom--
		}

		if left <= right {
			for i := bottom; i >= top; i-- {
				res = append(res, snailMap[i][left])
			}
			left++
		}
	}
	return res
}
