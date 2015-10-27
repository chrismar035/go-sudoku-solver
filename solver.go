package solver

func NewPuzzle(ints [81]int) Puzzle {
	return puzzleFromArray(ints)
}

func puzzleFromArray(ints [81]int) Puzzle {
	return Puzzle{Initial: ints, Solution: Solve(ints)}
}

func Solve(puzzle [81]int) [81]int {
	working := workingFromArray(puzzle)
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

	return working.toArray()
}
