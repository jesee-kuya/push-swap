package sortLarge

import "testing"

var TestCases2 = []struct {
	name       string
	lastLower  int
	lowerLimit int
	upperLimit int
}{
	{"Test1", 21, 21, 41},
}

func TestLimits(t *testing.T) {
	for _, tc := range TestCases2 {
		t.Run(tc.name, func(t *testing.T) {
			lowerLimit, upperLimit := Limits(tc.lastLower)

			if lowerLimit != tc.lowerLimit || upperLimit != tc.upperLimit {
				t.Errorf("The lowerlimit and upperlimit was %v and %v  but expected %v and %v\n", lowerLimit, upperLimit, tc.lowerLimit, tc.upperLimit)
			}
		})
	}
}
