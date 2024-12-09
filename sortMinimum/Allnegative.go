package sortMinimum

func AllNegative(arr []int) bool {
	for _, v := range arr {
		if v > 0 {
			return false
		}
	}
	return true
}
