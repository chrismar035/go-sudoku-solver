package solver

import "github.com/chrismar035/sudoku-solver/grid"

type multiBacktrackingSolver struct{}

func (b multiBacktrackingSolver) Solve(given Grid) []Grid {
	var puzzle [81]backtrackSquare
	for i, value := range given {
		puzzle[i] = backtrackSquare{value: value, initial: value != 0}
	}

	var solutions []Grid

	forward := true
	i := 0
	for {
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
			}
		}
		if i >= 81 {
			// Puzzle is done record this solution and go back and find more
			var ended [81]int
			for i, square := range puzzle {
				ended[i] = square.value
			}
			solutions = append(solutions, ended)

			forward = false
			i--
		}
		if i <= 0 {
			// Found all the solutions
			break
		}
	}

	return solutions
}
