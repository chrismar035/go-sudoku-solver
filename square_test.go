package solver

import (
	"testing"
)

func TestNewSquare(t *testing.T) {
	value := 5
	actual := newSquare(value)
	expected := square{Candidates: [9]bool{
		true,
		true,
		true,
		true,
		true,
		true,
		true,
		true,
		true},
		Value: value}

	if expected != actual {
		t.Errorf("newSquare(%d) == %v; want %v", value, actual, expected)
	}
}
