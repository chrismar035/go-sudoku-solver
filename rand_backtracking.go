package solver

import (
	"math/rand"
	"time"

	"github.com/chrismar035/sudoku-solver/grid"
)

type randBacktrackSquare struct {
	value          int
	initial        bool
	candidates     []int
	candidateIndex int
}

func (s randBacktrackSquare) outOfCandidates() bool {
	return s.candidateIndex >= len(s.candidates)
}

func (s randBacktrackSquare) nextCandidate() int {
	return s.candidates[s.candidateIndex]
}

type randBacktrackingSolver struct{}

func (b randBacktrackingSolver) Solve(given Grid) Grid {
	var puzzle [81]randBacktrackSquare
	for i, value := range given {
		puzzle[i] = randBacktrackSquare{
			value:          value,
			initial:        value != 0,
			candidates:     mixedValues(),
			candidateIndex: -1}
	}

	forward := true

	for i := 0; i < 81; {
		if puzzle[i].initial {
			if forward {
				i++
			} else {
				i--
			}
		} else {
			forward = true
			ok := false
			puzzle[i].candidateIndex++
			if puzzle[i].outOfCandidates() {
				puzzle[i].candidateIndex = -1
				puzzle[i].value = 0
				forward = false
				i--
			} else {
				puzzle[i].value = puzzle[i].nextCandidate()
				ok = randCheckValues(grid.IndicesForColumn(i), puzzle[i].value, puzzle)
				if ok {
					ok = randCheckValues(grid.IndicesForRow(i), puzzle[i].value, puzzle)
				}
				if ok {
					ok = randCheckValues(grid.IndicesForSub(i), puzzle[i].value, puzzle)
				}
			}
			if ok {
				i++
			}
		}
	}

	var ended Grid
	for i, square := range puzzle {
		ended[i] = square.value
	}
	return ended
}

func mixedValues() []int {
	rand.Seed(time.Now().UTC().UnixNano())
	ints := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	mixed := []int{}
	for len(ints) > 0 {
		i := rand.Int() % len(ints)
		mixed = append(mixed, ints[i])
		ints = append(ints[0:i], ints[i+1:]...)
	}

	return mixed
}

func randCheckValues(indices [8]int, current int, puzzle [81]randBacktrackSquare) bool {
	for _, index := range indices {
		if puzzle[index].value == current {
			return false
		}
	}
	return true
}
