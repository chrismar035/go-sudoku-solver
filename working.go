package solver

import "github.com/chrismar035/grid"

type working [81]square
type workingSet struct {
	Square square
	Index  int
}

func processSquare(i int, working working) working {
	square := working[i]

	square.checkIndices(grid.IndicesForRow(i), working)
	square.checkIndices(grid.IndicesForColumn(i), working)
	square.checkIndices(grid.IndicesForSub(i), working)

	square.tryToSetValueFromCandidates()

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

func workingFromArray(p [81]int) working {
	var working working
	for i, value := range p {
		working[i] = newSquare(value)
	}
	return working
}

func (w working) toArray() (array [81]int) {
	for i, square := range w {
		array[i] = square.Value
	}
	return
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
