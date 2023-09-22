package table_driven_math

import (
	"math"
	"testing"
)

type testCase struct {
	arg1     float64
	arg2     float64
	expected float64
}

var cases = []testCase{
	{
		arg1:     1.0,
		arg2:     2.0,
		expected: 2.0,
	},
	{
		arg1:     -100,
		arg2:     -200,
		expected: -100,
	},
	{
		arg1:     0,
		arg2:     -200,
		expected: 0,
	},
	{
		arg1:     -8.31373e-02,
		arg2:     1.84273e-02,
		expected: 1.84273e-02,
	},
}

func TestMathMax(t *testing.T) {
	for testIndex, tc := range cases {
		actual := math.Max(tc.arg1, tc.arg2)
		if actual != tc.expected {
			t.Errorf("Test %d: Max(%f, %f): Expected %f, got %f", testIndex, tc.arg1, tc.arg2, tc.expected, actual)
		}
	}
}
