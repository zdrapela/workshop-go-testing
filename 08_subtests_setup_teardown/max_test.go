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

func setupSubtest(t *testing.T) {
	t.Logf("[SETUP] Hello 👋!")
}

func teardownSubtest(t *testing.T) {
	t.Logf("[TEARDOWN] Bye, bye 🖖!")
}

func TestMax(t *testing.T) {
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			setupSubtest(t)
			defer teardownSubtest(t)
			RunTestMax(t, tc)
		})
	}
}

func RunTestMax(t *testing.T, tc struct {
	name           string
	input          []int
	expectedOutput int
}) {
	actual := Max(tc.input)

	t.Logf("[TEST] Hello from subtest %s \n", tc.name)

	if actual != tc.expectedOutput {
		t.Errorf("Expected %d, got %d", tc.expectedOutput, actual)
	}
}
