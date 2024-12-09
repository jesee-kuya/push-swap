package sortMinimum

import "testing"

var TestCases12 = []struct {
	name string
	arr  []int
	res  bool
}{
	{"Test1", []int{-1, -2, -3, -4}, true},
	{"Test2", []int{1, 2, 3, -4}, false},
}

func TestAllNegative(t *testing.T) {
	for _, tc := range TestCases12 {
		t.Run(tc.name, func(t *testing.T) {
			res := AllNegative(tc.arr)

			if res != tc.res {
				t.Errorf("Expected %v but got %v", tc.res, res)
			}
		})
	}
}
