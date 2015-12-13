package solver

type logicalSquare struct {
	Candidates [9]bool
	Value      int
}

type logicalSolver struct{}

func (b logicalSolver) Solve(given Grid) Grid {
	working := workingFromGrid(given)
	blankCount := working.blankCount()

	for {
		for i, square := range working {
			if square.Value == 0 {
				working = processSquare(i, working)
			}
		}
		newBlankCount := working.blankCount()
		if newBlankCount == blankCount {
			break
		} else {
			blankCount = newBlankCount
		}
	}

	return working.ToGrid()
}
