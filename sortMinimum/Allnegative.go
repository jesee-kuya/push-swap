package sortMinimum

// Checks if all elements are less than 0
func AllNegative(arr []int) bool {
	for _, v := range arr {
		if v > 0 {
			return false
		}
	}
	return true
}
