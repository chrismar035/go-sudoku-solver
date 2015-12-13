package solver

import "github.com/chrismar035/solver/grid"

type working [81]square
type workingSet struct {
	Square square
	Index  int
}

func processSquare(i int, working working) working {
	square := working[i]

	square.checkValues(grid.IndicesForRow(i), working)
	square.checkValues(grid.IndicesForColumn(i), working)
	square.checkValues(grid.IndicesForSub(i), working)


	if working[i].Value != square.Value {
		working[i] = square
		for _, index := range grid.IndicesForRow(i) {
			if working[index].Value == 0 {
				working = processSquare(index, working)
			}
		}
		for _, index := range grid.IndicesForColumn(i) {
			if working[index].Value == 0 {
				working = processSquare(index, working)
			}
		}
		for _, index := range grid.IndicesForSub(i) {
			if working[index].Value == 0 {
				working = processSquare(index, working)
			}
		}
	}
	return working
}

func workingFromGrid(p Grid) working {
	var working working
	for i, value := range p {
		working[i] = newSquare(value)
	}
	return working
}

func (w working) blankCount() int {
	count := 0
	for _, square := range w {
		if square.Value == 0 {
			count++
		}
	}
	return count
}

func (w working) ToGrid() (grid Grid) {
	for i, square := range w {
		grid[i] = square.Value
	}
	return
}
