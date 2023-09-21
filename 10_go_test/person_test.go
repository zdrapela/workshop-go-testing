package package_name

import (
	"testing"
)

func TestNewPersonPositiveAge(t *testing.T) {
	_, err := NewPerson(1)
	if err != nil {
		t.Errorf("Expected person, received %v", err)
	}
}

func TestNewPersonNegativeAge(t *testing.T) {
	p, err := NewPerson(-1)
	if err == nil {
		t.Errorf("Expected error, received %v", p)
	}
}

func TestOlderFirstOlderThanSecond(t *testing.T) {

	// if testing.Short() {
	// 	t.Skip("Skipping long-running test.")
	// }
	p1, _ := NewPerson(1)
	// p1, _ := NewPerson(100)
	p2, _ := NewPerson(2)

	if p1.older(p2) {
		t.Errorf("Expected p1 with age %d to be younger than p2 with age %d", p1.age, p2.age)
	}
}

func TestOlderSecondOlderThanFirst(t *testing.T) {
	p1, _ := NewPerson(2)
	p2, _ := NewPerson(1)

	if !p1.older(p2) {
		t.Errorf("Expected p1 with age %d to be older than p2 with age %d", p1.age, p2.age)
	}
}

// func TestNewPersonHugeAge(t *testing.T) {
// 	p, err := NewPerson(200)
// 	if err == nil {
// 		t.Errorf("Expected error, received %v", p)
// 	}
// }
