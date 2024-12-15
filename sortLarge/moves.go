package sortLarge

func MovesDown(upperLimit int, lowerLimit int, arr []int) (int, bool) {
	for i := 0; i < len(arr); i++ {
		if arr[i] >= lowerLimit && arr[i] <= upperLimit {
			return i + 1, true
		}
	}
	return 0, false
}

func MovesUp(upperLimit int, lowerLimit int, arr []int) (int, bool) {
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] >= lowerLimit && arr[i] <= upperLimit {
			return len(arr) - i, true
		}
	}
	return 0, false
}
