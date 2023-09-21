package subtests

import (
	"fmt"
	"os"
	"testing"
)

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

func TestMain(m *testing.M) {
	fmt.Println("[SETUP ALL TESTS] Hello 👋!")

	exitCode := m.Run()

	fmt.Println("[TEARDOWN ALL TESTS] Bye, bye 🖖!")

	os.Exit(exitCode)
}

// func TestMain "is" in every test file
// func TestMain(m *testing.M) {
// 	os.Exit(m.Run())
// }

func TestMax(t *testing.T) {
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
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
