package kata

func Grow(arr []int) int{
  prod := 1

	for _, value := range arr {
		prod *= value
	}

	return prod
  
}
