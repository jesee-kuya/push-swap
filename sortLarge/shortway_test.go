package sortLarge

import "testing"

var TestCases1 = []struct {
	name   string
	arr    []int
	check1 int
	check2 int
	res1   int
	res2   string
}{
	{"Test1", []int{34, 56, 24, 78, 88}, 20, 40, 34, "front"},
	{"Test2", []int{45, 24, 56, 46, 34, 50}, 20, 40, 24, "front"},
	{"Test3", []int{56, 57, 44, 24, 78}, 20, 40, 24, "back"},
}

func TestShortWay(t *testing.T) {
	for _, tc := range TestCases1 {
		t.Run(tc.name, func(t *testing.T) {
			res1, res2 := ShortWay(tc.arr, tc.check1, tc.check2)
			if res1 != tc.res1 || res2 != tc.res2 {
				t.Errorf("Expected the value %v and %v but got %v and %v", tc.res1, tc.res2, res1, res2)
			}
		})
	}
}
