package testing

import "testing"

func TestMax(t *testing.T) {
	actual := Max([]int{1, 2, 3, 4})
	if actual != 4 {
		t.Errorf("Expected %d, got %d", 4, actual)
	}
}

// Every test file ends with *_test.go
// Tests are usually in
