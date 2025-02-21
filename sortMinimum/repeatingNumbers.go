package sortMinimum

func HasRepeatingNumbers(arr []int) bool {
	numMap := make(map[int]bool)

	for _, num := range arr {
		if numMap[num] {
			return true
		}
		numMap[num] = true
	}

	return false
}
