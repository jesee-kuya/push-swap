package sortMinimum

func Checker(input []int, ops []string) string {
	arr := append([]int{}, input...)
	var arrB []int

	for i := 0; i < len(ops); i++ {
		switch ops[i] {
		case "sa":
			if len(arr) > 1 {
				arr[0], arr[1] = arr[1], arr[0]
			} else {
				return "KO"
			}
		case "sb":
			if len(arrB) > 1 {
				arrB[0], arrB[1] = arrB[1], arrB[0]
			} else {
				return "KO"
			}
		case "ss":
			if len(arrB) > 1 {
				arrB[0], arrB[1] = arrB[1], arrB[0]
			} else {
				return "KO"
			}
			if len(arr) > 1 {
				arr[0], arr[1] = arr[1], arr[0]
			} else {
				return "KO"
			}
		case "pb":
			if len(arr) >= 1 {
				arrB = append([]int{arr[0]}, arrB...)
				arr = arr[1:]
			} else {
				return "KO"
			}
		case "pa":
			if len(arrB) >= 1 {
				arr = append([]int{arrB[0]}, arr...)
				arrB = arrB[1:]
			} else {
				return "KO"
			}
		case "ra":
			if len(arr) >= 1 {
				arr = RotateLeft(arr)
			} else {
				return "KO"
			}
		case "rb":
			if len(arrB) >= 1 {
				arrB = RotateLeft(arrB)
			} else {
				return "KO"
			}
		case "rr":
			if len(arrB) >= 1 {
				arrB = RotateLeft(arrB)
			} else {
				return "KO"
			}
			if len(arr) >= 1 {
				arr = RotateLeft(arr)
			} else {
				return "KO"
			}
		case "rrb":
			if len(arrB) >= 1 {
				arrB = RotateRight(arrB)
			} else {
				return "KO"
			}
		case "rra":
			if len(arr) >= 1 {
				arr = RotateRight(arr)
			} else {
				return "KO"
			}
		case "rrr":
			if len(arr) >= 1 {
				arr = RotateRight(arr)
			} else {
				return "KO"
			}
			if len(arrB) >= 1 {
				arrB = RotateRight(arrB)
			} else {
				return "KO"
			}
		}
	}
	if IsSorted(arr) && len(arrB) == 0 {
		return "OK"
	}
	return "KO"
}
