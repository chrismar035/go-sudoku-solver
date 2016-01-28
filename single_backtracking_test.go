package solver

import (
	"errors"
	"testing"
)

func TestSingleBacktrackingSolver(t *testing.T) {
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
		t.Errorf("singleBacktracking.Solve(%d) == %v; want %v", puzzle, actual, expected)
	}
}

func TestSingleBacktrackingSolverToo(t *testing.T) {
	puzzle := Grid{
		6, 2, 5, 9, 1, 3, 7, 8, 4,
		7, 9, 3, 2, 8, 4, 6, 5, 1,
		1, 8, 4, 6, 7, 5, 9, 2, 3,
		9, 7, 1, 5, 2, 8, 3, 4, 6,
		3, 5, 6, 4, 9, 1, 8, 7, 2,
		2, 4, 8, 7, 3, 6, 5, 1, 9,
		5, 1, 2, 8, 6, 9, 4, 3, 7,
		4, 6, 7, 0, 5, 2, 1, 9, 8,
		8, 3, 9, 1, 4, 7, 2, 6, 5,
	}
	expected := Grid{
		6, 2, 5, 9, 1, 3, 7, 8, 4,
		7, 9, 3, 2, 8, 4, 6, 5, 1,
		1, 8, 4, 6, 7, 5, 9, 2, 3,
		9, 7, 1, 5, 2, 8, 3, 4, 6,
		3, 5, 6, 4, 9, 1, 8, 7, 2,
		2, 4, 8, 7, 3, 6, 5, 1, 9,
		5, 1, 2, 8, 6, 9, 4, 3, 7,
		4, 6, 7, 3, 5, 2, 1, 9, 8,
		8, 3, 9, 1, 4, 7, 2, 6, 5,
	}

	backtracker := NewBacktrackingSolver()
	actual, _ := backtracker.Solve(puzzle)
	if actual != expected {
		t.Errorf("singleBacktracking.Solve(%d) == %v; want %v", puzzle, actual, expected)
	}
}

func TestMultiSingleBacktrackingSolver(t *testing.T) {
	puzzle := Grid{
		7, 5, 0, 0, 0, 0, 0, 2, 0,
		1, 0, 0, 2, 0, 0, 0, 0, 0,
		3, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 1, 7, 0, 0, 0, 0,
		0, 0, 1, 0, 3, 0, 5, 0, 0,
		0, 0, 0, 0, 4, 8, 0, 0, 0,
		8, 0, 9, 0, 5, 0, 0, 0, 2,
		0, 0, 0, 0, 0, 7, 0, 0, 3,
		0, 6, 0, 0, 0, 0, 0, 5, 1,
	}

	backtracker := NewBacktrackingSolver()
	_, err := backtracker.Solve(puzzle)
	expected := errors.New("Multiple solutions found")
	if err != nil && err.Error() != expected.Error() {
		t.Errorf("singleBacktracking.Solve(%d) == %v; want %v", puzzle, err, expected)
	}
}
