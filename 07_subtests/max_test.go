package subtests

import "testing"

var cases = []struct {
	name           string
	input          []int
	expectedOutput int
}{
	{
		name:           "FourNumbersFourthNumberMax",
		input:          []int{1, 2, 3, 4},
		expectedOutput: 4,
	},
	{
		name:           "ThreeNumbersFirstNumberMax",
		input:          []int{9, 2, 3},
		expectedOutput: 9,
	},
	{
		name:           "ZeroNumbers",
		input:          []int{},
		expectedOutput: -1,
	},
}

func TestMax(t *testing.T) {
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := Max(tc.input)

			if actual != tc.expectedOutput {
				t.Errorf("Expected %d, got %d", tc.expectedOutput, actual)
			}
		})
	}
}
