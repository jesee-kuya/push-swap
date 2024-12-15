package sortLarge

import "testing"

var TestCases3 = []struct {
	name       string
	lowerlimit int
	upperlimit int
	arr        []int
	moves      int
	check      bool
}{
	{"Test1", 20, 40, []int{1, 2, 3, 4, 56, 23, 41, 45}, 6, true},
}

var TestCases4 = []struct {
	name       string
	lowerlimit int
	upperlimit int
	arr        []int
	moves      int
	check      bool
}{
	{"Test1", 20, 40, []int{1, 2, 3, 4, 56, 23, 41, 45}, 3, true},
}

func TestMovesDown(t *testing.T) {
	for _, tc := range TestCases3 {
		t.Run(tc.name, func(t *testing.T) {
			moves, check := MovesDown(tc.upperlimit, tc.lowerlimit, tc.arr)

			if moves != tc.moves || check != tc.check {
				t.Errorf("Expected %v moves and the check to be %v\nbut got %v moves and a %v check\n", tc.moves, tc.check, moves, check)
			}
		})
	}
}

func TestMovesUp(t *testing.T) {
	for _, tc := range TestCases4 {
		t.Run(tc.name, func(t *testing.T) {
			moves, check := MovesUp(tc.upperlimit, tc.lowerlimit, tc.arr)

			if moves != tc.moves || check != tc.check {
				t.Errorf("Expected %v moves and the check to be %v\nbut got %v moves and a %v check\n", tc.moves, tc.check, moves, check)
			}
		})
	}
}
