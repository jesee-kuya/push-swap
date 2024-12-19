package sortMinimum

import (
	"reflect"
	"testing"
)

var TestCases17 = []struct {
	name      string
	moves     int
	turn      string
	tempstack []int
	expected  []int
	ops       []string
}{
	{"Test1", 3, "reverse", []int{20, 15, 14, 11, 56, 34, 25, 24, 23}, []int{56, 34, 25, 24, 23, 20, 15, 14, 11}, []string{"rb", "rb", "rb", "rb"}},
	{"Test2", 2, "rotate", []int{5, 3, 2, 1, 7, 6}, []int{7, 6, 5, 3, 2, 1}, []string{"rrb", "rrb"}},
}

func TestSortTemp(t *testing.T) {
	for _, tc := range TestCases17 {
		t.Run(tc.name, func(t *testing.T) {
			got, ops := SortTemp(tc.moves, tc.turn, tc.tempstack)

			if !reflect.DeepEqual(got, tc.expected) || !reflect.DeepEqual(ops, tc.ops) {
				t.Errorf("Expected the slice %v with %v  operations but got\n%v and %v operations", tc.expected, tc.ops, got, ops)
			}
		})
	}
}
