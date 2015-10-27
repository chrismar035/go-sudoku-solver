package grid

func IndicesForRow(rawIndex int) [8]int {
	row := rawIndex / 9
	var indices [8]int
	currentIndex := 0
	for i := 0; currentIndex < 8; i++ {
		index := 9*row + i
		if index != rawIndex {
			indices[currentIndex] = index
			currentIndex++
		}
	}
	return indices
}

func IndicesForColumn(rawIndex int) [8]int {
	column := rawIndex % 9
	var indices [8]int
	currentIndex := 0
	for i := 0; currentIndex < 8; i++ {
		index := i*9 + column
		if index != rawIndex {
			indices[currentIndex] = index
			currentIndex++
		}
	}
	return indices
}

func IndicesForSub(rawIndex int) [8]int {
	row := rawIndex / 9
	rowOffSet := row / 3
	column := rawIndex % 9
	columnOffset := column / 3
	baseIndices := [9]int{0, 1, 2, 9, 10, 11, 18, 19, 20}

	var indices [8]int
	currentIndex := 0
	for i := 0; currentIndex < 8; i++ {
		value := baseIndices[i]
		candidate := value + (3 * columnOffset) + (27 * rowOffSet)
		if candidate != rawIndex {
			indices[currentIndex] = candidate
			currentIndex++
		}
	}

	return indices
}
