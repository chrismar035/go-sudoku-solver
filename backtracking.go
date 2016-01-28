package solver

import (
	"github.com/chrismar035/sudoku-solver/grid"
)

type backtrackSquare struct {
	value   int
	initial bool
}

type backtrackingSolver struct{}

func (b backtrackingSolver) Solve(given Grid) (Grid, error) {
	var puzzle [81]backtrackSquare
	for i, value := range given {
		puzzle[i] = backtrackSquare{value: value, initial: value != 0}
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
			puzzle[i].value++
			ok := false
			if puzzle[i].value > 9 {
				puzzle[i].value = 0
				forward = false
				i--
			} else {
				ok = checkValues(grid.IndicesForColumn(i), puzzle[i].value, puzzle)
				if ok {
					ok = checkValues(grid.IndicesForRow(i), puzzle[i].value, puzzle)
				}
				if ok {
					ok = checkValues(grid.IndicesForSub(i), puzzle[i].value, puzzle)
				}
			}
			if ok {
				i++
			} else {
			}
		}
		// time.Sleep(time.Second)
	}

	var ended [81]int
	for i, square := range puzzle {
		ended[i] = square.value
	}
	return ended, nil
}

func checkValues(indices [8]int, current int, puzzle [81]backtrackSquare) bool {
	for _, index := range indices {
		if puzzle[index].value == current {
			return false
		}
	}
	return true
}
