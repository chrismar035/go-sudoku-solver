package backtracker

import "github.com/chrismar035/solver/grid"

type square struct {
	value   int
	initial bool
}

func Solve(given [81]int) [81]int {
	var puzzle [81]square
	for i, value := range given {
		puzzle[i] = square{value: value, initial: value != 0}
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
	return ended
}

func checkValues(indices [8]int, current int, puzzle [81]square) bool {
	for _, index := range indices {
		if puzzle[index].value == current {
			return false
		}
	}
	return true
}
