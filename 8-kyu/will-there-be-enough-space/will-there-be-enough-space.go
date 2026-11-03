package kata

func Enough (cap, on, wait int) int {
  freeSpace := cap - on

	if freeSpace >= wait {
		return 0
	}

	return wait - freeSpace
}
