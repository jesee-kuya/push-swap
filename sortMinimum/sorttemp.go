package sortMinimum

func SortTemp(moves int, turn string, tempstack []int) ([]int, []string) {
	var ops []string
	if turn == "reverse" {
		moves += 1
		for moves > 0 {
			tempstack = RotateLeft(tempstack)
			ops = append(ops, "rb")
			moves--
		}
	}
	if turn == "rotate" {
		for moves > 0 {
			tempstack = RotateRight(tempstack)
			ops = append(ops, "rrb")
			moves--
		}
	}
	return tempstack, ops
}
