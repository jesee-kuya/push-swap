package sortMinimum

import "testing"

var TestCases18 = []struct {
	name     string
	arr      []int
	ops      []string
	expected string
}{
	{"Test1", []int{2, 1, 3, 6, 5, 8}, []string{"sa", "pb", "pb", "pb", "sa", "pa", "pa", "pa"}, "OK"},
	{"Test2", []int{0, 9, 1, 8, 2}, []string{"pb", "ra", "pb", "ra", "sa", "ra", "pa", "pa"}, "OK"},
	{"Test3", []int{924, 169, 992, 451, 685}, []string{"sa", "pb", "rra", "rra", "pa"}, "OK"},
	{"Test4", []int{9, 10, 1, 4, 15}, []string{"sa", "sa", "pb", "pb", "sa"}, "KO"},
}

func TestChecker(t *testing.T) {
	for _, tc := range TestCases18 {
		t.Run(tc.name, func(t *testing.T) {
			got := Checker(tc.arr, tc.ops)

			if got != tc.expected {
				t.Errorf("Expected %v but got %v", tc.expected, got)
			}
		})
	}
}
