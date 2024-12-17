package sortMinimum

func RangeDown(minVal int, maxVal int, arr []int) (int, int) {
	for i := 0; i < len(arr); i++ {
		if arr[i] >= minVal && arr[i] <= maxVal {
			return arr[i], i
		}
	}
	return 0, -1
}

func RangeUp(minVal int, maxVal int, arr []int) (int, int) {
	for i := len(arr) - 1; i >= 0 ; i-- {
		if arr[i] >= minVal && arr[i] <= maxVal {
			return arr[i], len(arr) - i
		}
	}
	return 0, -1
}
