package sortMinimum

func PreSort(num int, tempStack []int) ([]int, string, []string, int) {
	var ind int
	var ops []string
	var turn string
	var moves int
	for i := 0; i < len(tempStack); i++ {
		if tempStack[i] < num {
			ind = i
			break
		}
	}
	checker := (len(tempStack) - 1) - ind

	if checker+1 < ind {
		moves = checker
		checker++
		for checker > 0 {
			tempStack = RotateRight(tempStack)
			ops = append(ops, "rrb")
			checker--
		}
		turn = "reverse"

	}
	if checker+1 > ind {
		moves = ind
		for ind > 0 {
			tempStack = RotateLeft(tempStack)
			ops = append(ops, "rb")
			ind--
		}
		turn = "rotate"
	}
	if checker == ind {
		moves = checker
		for checker > 0 {
			tempStack = RotateRight(tempStack)
			ops = append(ops, "rrb")
			checker--
		}
		turn = "reverse"
	}
	return tempStack, turn, ops, moves
}
