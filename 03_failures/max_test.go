package failures

import "testing"

func TestMax(t *testing.T) {
	actual := Max([]int{1, 2, 3, 4})
	if actual != 4 {
		t.Errorf("Expected %d, got %d", 4, actual)
	}
}

func TestMaxInvalid(t *testing.T) {
	actual := Max([]int{1, 2, 3, 4})
	if actual != 5 {
		t.Errorf("Expected %d, got %d", 5, actual)
	}
}

func TestMaxEmpty(t *testing.T) {
	actual := Max([]int{})
	if actual != -1 {
		t.Errorf("Expected %v, got %d", -1, actual)
	}
}
