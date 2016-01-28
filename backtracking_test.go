package solver

import (
	"errors"
	"testing"
)

func TestBacktrackingSolver(t *testing.T) {
	puzzle := Grid{
		7, 5, 0, 0, 0, 0, 0, 2, 0,
		1, 0, 0, 2, 0, 0, 0, 0, 0,
		3, 0, 0, 0, 9, 0, 4, 0, 6,
		0, 0, 0, 1, 7, 0, 0, 0, 0,
		0, 0, 1, 0, 3, 0, 5, 0, 0,
		0, 0, 0, 0, 4, 8, 0, 0, 0,
		8, 0, 9, 0, 5, 0, 0, 0, 2,
		0, 0, 0, 0, 0, 7, 0, 0, 3,
		0, 6, 0, 0, 0, 0, 0, 5, 1,
	}
	expected := Grid{
		7, 5, 4, 6, 1, 3, 9, 2, 8,
		1, 9, 6, 2, 8, 4, 3, 7, 5,
		3, 2, 8, 7, 9, 5, 4, 1, 6,
		9, 8, 5, 1, 7, 6, 2, 3, 4,
		6, 4, 1, 9, 3, 2, 5, 8, 7,
		2, 3, 7, 5, 4, 8, 1, 6, 9,
		8, 7, 9, 3, 5, 1, 6, 4, 2,
		5, 1, 2, 4, 6, 7, 8, 9, 3,
		4, 6, 3, 8, 2, 9, 7, 5, 1,
	}

	backtracker := NewBacktrackingSolver()
	actual, _ := backtracker.Solve(puzzle)
	if actual != expected {
		t.Errorf("backtracking.Solve(%d) == %v; want %v", puzzle, actual, expected)
	}
}

func TestUnsolveableBacktrackingSolver(t *testing.T) {
	puzzle := Grid{
		1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 9,
		1, 2, 3, 4, 5, 6, 7, 8, 0,
	}

	backtracker := NewBacktrackingSolver()
	_, err := backtracker.Solve(puzzle)
	expected := errors.New("Unsolvable puzzle")
	if err != nil && err.Error() != expected.Error() {
		t.Errorf("backtracking.Solve(%d) == %v; want %v", puzzle, err, expected)
	}
}
