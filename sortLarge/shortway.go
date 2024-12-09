package sortLarge

func ShortWay(arr []int, check1 int, check2 int) (int, string) {
	var val1, val2, ind1, ind2 int

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] >= check1 && arr[i] <= check2 {
			val1 = arr[i]
			ind1 = i
			break
		}
	}

	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] >= check1 && arr[i] <= check2 {
			val2 = arr[i]
			ind2 = (len(arr)-1) - i
			break
		}
	}

	if ind1 > ind2 {
		return val2, "back"
	} 
	if ind1 < ind2 {
		return val1, "front"
	}
	if ind1 == ind2 {
		return val1, "front"
	}
	return 0, ""
}
