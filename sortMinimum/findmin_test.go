package sortMinimum

import (
	"testing"
)

var TestCases13 = []struct {
	name     string
	arr      []int
	expected int
}{
	{"Test1", []int{7, 9, 8, 5, 3, 4, 6}, 3},
}

func TestFindMin(t *testing.T) {
	for _, tc := range TestCases13 {
		t.Run(tc.name, func(t *testing.T) {
			got := FindMin(tc.arr)

			if got != tc.expected {
				t.Errorf("Expected %v but got %v", got, tc.expected)
			}
		})
	}
}
