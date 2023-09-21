//go:build example_tag
// +build example_tag

// the second one is a legacy style of tag

package go_test_tag

import (
	"testing"
)

func TestNewPersonPositiveAge(t *testing.T) {
	_, err := NewPerson(1)
	if err != nil {
		t.Errorf("Expected person, received %v", err)
	}
}
