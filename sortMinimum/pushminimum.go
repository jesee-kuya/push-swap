package sortMinimum

// PushMinimum finds the sequence of operations to sort the array.
func PushMinimum(input []int) []string {
	operations := []string{}
	arr := append([]int{}, input...)
	count := 0

	if len(arr) > 9 {
		for !IsSorted(arr) {
			ops, updatedArr, tempStack := PreEmpt(arr)
			operations = append(operations, ops...)
			resl, finalArr := ProcessTempStack(updatedArr, tempStack)
			operations = append(operations, resl...)
			arr = finalArr
			count++
		}
	} else {
		for !IsSorted(arr) {
			ops, updatedArr, tempStack := PushSwap(arr)
			operations = append(operations, ops...)
			resl, finalArr := ProcessTempStack(updatedArr, tempStack)
			operations = append(operations, resl...)
			arr = finalArr
		}
	}

	operations = SortOperations(operations)
	return operations
}
