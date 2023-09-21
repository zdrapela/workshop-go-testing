package go_test_tag

import "errors"

type Person struct {
	age int
}

func NewPerson(age int) (*Person, error) {
	if age < 1 {
		return nil, errors.New("a person is at least 1 years old")
	}

	// if age >= 150 {
	// 	return nil, errors.New("a person cannot be older than 150 years")
	// }

	return &Person{
		age: age,
	}, nil
}

func (p *Person) older(other *Person) bool {
	return p.age > other.age
}
