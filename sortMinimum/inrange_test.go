package sortMinimum

import "testing"

var TestCases14 = []struct {
	name     string
	minVal   int
	maxVal   int
	arr      []int
	expected int
	moves    int
}{
	{"Test1", 5, 25, []int{89, 67, 56, 45, 34, 27, 23, 21}, 23, 6},
}

var TestCases15 = []struct {
	name     string
	minVal   int
	maxVal   int
	arr      []int
	expected int
	moves    int
}{
	{"Test1", 5, 25, []int{89, 67, 56, 45, 34, 27, 23, 21}, 21, 1},
}

func TestRangeDown(t *testing.T) {
	for _, tc := range TestCases14 {
		t.Run(tc.name, func(t *testing.T) {
			got, moves := RangeDown(tc.minVal, tc.maxVal, tc.arr)
			if got != tc.expected || moves != tc.moves {
				t.Errorf("Expected %v in %v moves but got the %v in %v moves", tc.expected, tc.moves, got, moves)
			}
		})
	}
}

func TestRangeUp(t *testing.T) {
	for _, tc := range TestCases15 {
		t.Run(tc.name, func(t *testing.T) {
			got, moves := RangeUp(tc.minVal, tc.maxVal, tc.arr)
			if got != tc.expected || moves != tc.moves {
				t.Errorf("Expected %v in %v moves but got the %v in %v moves", tc.expected, tc.moves, got, moves)
			}
		})
	}
}
