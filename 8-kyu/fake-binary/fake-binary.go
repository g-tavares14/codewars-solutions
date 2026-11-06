package kata

func FakeBin(x string) string {
	var binaryString string

	for _, char := range x {
		if char < '5' {
			binaryString += "0"
		} else {
			binaryString += "1"
		}
	}
	return binaryString
}
