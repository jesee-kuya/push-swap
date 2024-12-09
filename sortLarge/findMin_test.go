package sortLarge

import "testing"

var TestCases = []struct {
	name     string
	input    []int
	minVal   int
	minIndex int
}{
	{"Test1", []int{55, 3, 98, 400}, 3, 1},
	{"Test2", []int{44, 678, 490, 708}, 44, 0},
}

func TestFindMin(t *testing.T) {
	for _, tc := range TestCases {
		t.Run(tc.name, func(t *testing.T) {
			minVal, minIndex := FindMin(tc.input)
			if minIndex != tc.minIndex || minVal != tc.minVal {
				t.Errorf("Expected the value %v and index %v but got the value %v and index %v", tc.minVal, tc.minIndex, minVal, minIndex)
			}
		})
	}
}
