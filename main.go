package main

import (
	"fmt"
	"sync"
)

func main() {
	puzzle := Puzzle{Initial: [81]int{0, 1, 0, 0, 0, 6, 5, 2, 7,
		7, 8, 0, 1, 4, 5, 0, 0, 9,
		0, 0, 0, 0, 2, 0, 0, 1, 0,

		0, 0, 5, 0, 0, 0, 7, 4, 6,
		0, 0, 0, 9, 0, 7, 0, 0, 0,
		6, 7, 1, 0, 0, 0, 9, 0, 0,

		0, 3, 0, 0, 9, 0, 0, 0, 0,
		9, 0, 0, 4, 8, 3, 0, 6, 5,
		1, 6, 8, 5, 0, 0, 0, 9, 0}}

	var wg sync.WaitGroup
	working := WorkingFromPuzzle(puzzle)
	newSets := make(chan WorkingSet)
	processSquare := func(i int) {
		defer wg.Done()
		square := working[i]

		// fmt.Println("Processing on:", r, "-", c)
		// mark same row as non-candidates
		for _, index := range indicesForRow(i) {
			neighbor := working[index]
			if neighbor.Value != 0 {
				square.Candidates[neighbor.Value-1] = false
			}
		}

		// mark same column as non-candidates
		for _, index := range indicesForColumn(i) {
			neighbor := working[index]
			if neighbor.Value != 0 {
				square.Candidates[neighbor.Value-1] = false
			}
		}

		// mark same sub-square as non-candidates
		for _, index := range indicesForSub(i) {
			neighbor := working[index]
			if neighbor.Value != 0 {
				square.Candidates[neighbor.Value-1] = false
			}
		}

		var candidate int
		candidateCount := 0
		for i, candidateCheck := range square.Candidates {
			if candidateCheck {
				candidate = i
				candidateCount++
			}
		}
		if candidateCount == 1 {
			square.Value = candidate + 1
		}

		newSets <- WorkingSet{Square: square, Index: i}
	}
	for i, square := range working {
		if square.Value == 0 {
			wg.Add(1)
			go processSquare(i)
		}
	}

	go func() {
		wg.Wait()
		close(newSets)
	}()

	for newSet := range newSets {
		index := newSet.Index
		row := index / 9
		column := index % 9
		square := newSet.Square
		// fmt.Println("Received:", row, "-", column, square)
		//fmt.Println("Comparing:", working[column][row].Value, square.Value)
		if working[index].Value != square.Value {
			fmt.Println("Got value at", row, "-", column, "of", square.Value)
			working[index] = square
			for _, index := range indicesForRow(index) {
				if working[index].Value == 0 {
					wg.Add(1)
					go processSquare(index)
				}
			}
			for _, index := range indicesForColumn(index) {
				if working[index].Value == 0 {
					wg.Add(1)
					go processSquare(index)
				}
			}
			for _, index := range indicesForSub(index) {
				if working[index].Value == 0 {
					wg.Add(1)
					go processSquare(index)
				}
			}
		}
	}

	fmt.Println("Worked on:")
	fmt.Println(puzzle)
	fmt.Println()
	fmt.Println("Solution?")
	fmt.Println(PuzzleFromWorking(working))
}

func indicesForRow(raw_index int) [8]int {
	row := raw_index / 9
	var indices [8]int
	for i := 0; i < 8; i++ {
		index := 9*row + i
		if index != raw_index {
			indices[i] = index
		}
	}
	return indices
}

func indicesForColumn(raw_index int) [8]int {
	column := raw_index % 9
	var indices [8]int
	for i := 0; i < 8; i++ {
		index := i*9 + column
		if index != raw_index {
			indices[i] = index
		}
	}
	return indices
}

func indicesForSub(raw_index int) [8]int {
	row := raw_index / 9
	column := raw_index % 9
	base_indices := [9]int{0, 1, 2, 9, 10, 11, 18, 19, 20}

	var indices [8]int
	current_index := 0
	for i := 0; current_index < 8; i++ {
		value := base_indices[i]
		candidate := value + column + (3 * row)
		if candidate != raw_index {
			indices[current_index] = candidate
			current_index++
		}
	}

	return indices
}
