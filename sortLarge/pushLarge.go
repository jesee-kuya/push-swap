package sortLarge

import (
	min "push/sortMinimum"
)

func PushLarge(input []int) []string {
	var operations []string
	// var tempStack []int
	arr := append([]int{}, input...)
	// maxval, _ := min.FindMax(arr)
	minVal, _ := FindMin(arr)
	lowerLimit, upperLimit := Limits(minVal)

	for arr != nil {
		upMoves, check1 := MovesUp(upperLimit, lowerLimit, arr)
		downMoves, check2 := MovesUp(upperLimit, lowerLimit, arr)

		if check1 && check2 {
			if upMoves < downMoves {
				for upMoves > 0 {
					operations = append(operations, "rra")
					min.RotateRight(arr)
					upMoves--
				}
			}else {
				for downMoves > 0 {

				}
			}
		}
	}
	return operations
}
