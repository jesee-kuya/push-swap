package sortLarge

// FindMin locates the minimum value and its index in the array.
func FindMin(arr []int) (int, int) {
	minVal, minIndex := arr[0], 0
	for i, val := range arr {
		if val < minVal {
			minVal = val
			minIndex = i
		}
	}
	return minVal, minIndex
}
