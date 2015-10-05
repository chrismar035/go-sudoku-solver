package main

type Working [81]Square
type WorkingSet struct {
	Square Square
	Index  int
}

func WorkingFromPuzzle(p Puzzle) Working {
	var working Working
	for i, value := range p.Initial {
		working[i] = NewSquare(value)
	}
	return working
}

func PuzzleFromWorking(working Working) Puzzle {
	var puzzle Puzzle
	for i, square := range working {
		puzzle.Initial[i] = square.Value
	}
	return puzzle
}
