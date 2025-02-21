package sortMinimum

import "testing"

var TestCases19 = []struct {
	name string
	arr  []int
	res  bool
}{
	{"Test1", []int{1, 2, 3, 4, 5}, false},
	{"Test2", []int{1, 2, 3, 4, 5, 5}, true},
	{"Test3", []int{1, 2, 3, 4, 5, 5, 5}, true},
	{"Test4", []int{1, 2, 3, 4, 5, 5, 5, 5}, true},
	{"Test5", []int{1, 2, 3, 4, 5, 5, 5, 5, 5}, true},
}

func TestHasRepeatingNumbers(t *testing.T) {
	for _, tc := range TestCases19 {
		t.Run(tc.name, func(t *testing.T) {
			res := HasRepeatingNumbers(tc.arr)

			if res != tc.res {
				t.Errorf("Expected %v but got %v", tc.res, res)
			}
		})
	}
}
