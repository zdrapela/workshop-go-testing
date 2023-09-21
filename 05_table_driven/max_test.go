package table_driven

import "testing"

type testCase struct {
	input          []int
	expectedOutput int
}

var cases = []testCase{
	{input: []int{1, 2, 3, 4}, expectedOutput: 4},
	{expectedOutput: 3, input: []int{1, 2, 3, 2}},
	{[]int{}, -1},
	{
		input:          []int{0, 0, 0, 0},
		expectedOutput: 0,
	},
}

func TestMax(t *testing.T) {
	for _, tc := range cases {
		actual := Max(tc.input)
		if actual != tc.expectedOutput {
			t.Errorf("Expected %d, got %d", tc.expectedOutput, actual)
		}
	}
}
