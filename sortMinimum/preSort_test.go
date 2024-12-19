package sortMinimum

import (
	"reflect"
	"testing"
)

var TestCases16 = []struct {
	name      string
	num       int
	tempStack []int
	expected  []int
	turn      string
	ops       []string
	moves     int
}{
	{"Test1", 20, []int{56, 34, 25, 24, 23, 15, 14, 11}, []int{15, 14, 11, 56, 34, 25, 24, 23}, "reverse", []string{"rrb", "rrb", "rrb"}, 2},
	{"Test2", 5, []int{7, 6, 3, 2, 1}, []int{3, 2, 1, 7, 6}, "rotate", []string{"rb", "rb"}, 2},
}

func TestPreSort(t *testing.T) {
	for _, tc := range TestCases16 {
		t.Run(tc.name, func(t *testing.T) {
			got, turn, ops, moves := PreSort(tc.num, tc.tempStack)

			if !reflect.DeepEqual(got, tc.expected) || turn != tc.turn || !reflect.DeepEqual(ops, tc.ops) || moves != tc.moves {
				t.Errorf("Expected %v\nBut got %v", tc.expected, got)
			}
		})
	}
}
