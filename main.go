package main

import (
	"bytes"
	"fmt"
	"strconv"
	"sync"
)

type Row [9]int
type Rows [9]Row

type Puzzle struct {
	Initial  Rows
	Solution Rows
}

type WorkingRow [9]Square
type Working [9]WorkingRow
type WorkingSet struct {
	Square Square
	Row    int
	Column int
}

type Square struct {
	Candidates [9]bool
	Value      int
}

func main() {
	puzzle := Puzzle{Initial: Rows{
		Row{0, 1, 0, 0, 0, 6, 5, 2, 7},
		Row{7, 8, 0, 1, 4, 5, 0, 0, 9},
		Row{0, 0, 0, 0, 2, 0, 0, 1, 0},

		Row{0, 0, 5, 0, 0, 0, 7, 4, 6},
		Row{0, 0, 0, 9, 0, 7, 0, 0, 0},
		Row{6, 7, 1, 0, 0, 0, 9, 0, 0},

		Row{0, 3, 0, 0, 9, 0, 0, 0, 0},
		Row{9, 0, 0, 4, 8, 3, 0, 6, 5},
		Row{1, 6, 8, 5, 0, 0, 0, 9, 0}},

		Solution: Rows{
			Row{0, 1, 0, 0, 0, 6, 5, 2, 7},
			Row{7, 8, 0, 1, 4, 5, 0, 0, 9},
			Row{0, 0, 0, 0, 2, 0, 0, 1, 0},

			Row{0, 0, 5, 0, 0, 0, 7, 4, 6},
			Row{0, 0, 0, 9, 0, 7, 0, 0, 0},
			Row{6, 7, 1, 0, 0, 0, 9, 0, 0},

			Row{0, 3, 0, 0, 9, 0, 0, 0, 0},
			Row{9, 0, 0, 4, 8, 3, 0, 6, 5},
			Row{1, 6, 8, 5, 0, 0, 0, 9, 0}}}

	var wg sync.WaitGroup
	working := WorkingFromPuzzle(puzzle)
	newSets := make(chan WorkingSet)
	processSquare := func(r, c int) {
		defer wg.Done()
		square := working[r][c]

		// fmt.Println("Processing on:", r, "-", c)
		// mark same row as non-candidates
		for _, neighbor := range working[r] {
			if neighbor.Value != 0 {
				square.Candidates[neighbor.Value-1] = false
			}
		}

		// mark same column as non-candidates
		for _, row := range working {
			neighbor := row[c]
			if neighbor.Value != 0 {
				square.Candidates[neighbor.Value-1] = false
			}
		}

		// mark same sub-square as non-candidates
		// TODO

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

		newSets <- WorkingSet{Square: square, Row: r, Column: c}
	}
	for i, workingRow := range working {
		for ii, square := range workingRow {
			if square.Value == 0 {
				wg.Add(1)
				go processSquare(i, ii)
			}
		}
	}

	go func() {
		wg.Wait()
		close(newSets)
	}()

	for newSet := range newSets {
		row := newSet.Row
		column := newSet.Column
		square := newSet.Square
		// fmt.Println("Received:", row, "-", column, square)
		//fmt.Println("Comparing:", working[column][row].Value, square.Value)
		if working[row][column].Value != square.Value {
			fmt.Println("Got value at", row, "-", column, "of", square.Value)
			working[row][column] = square
			for i := 0; i < 9; i++ {
				if i != row && working[row][i].Value == 0 {
					wg.Add(1)
					go processSquare(row, i)
				}
				if i != column && working[i][column].Value == 0 {
					wg.Add(1)
					go processSquare(i, column)
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

func (p Puzzle) String() string {
	var buffer bytes.Buffer
	for i, row := range p.Initial {
		if i != 0 && i%3 == 0 {
			buffer.WriteString("\n")
		}
		buffer.WriteString(row.String())
		buffer.WriteString("\n")
	}
	return buffer.String()
}

func (r Row) String() string {
	var buffer bytes.Buffer
	for i, value := range r {
		if i != 0 && i%3 == 0 {
			buffer.WriteString(" ")
		}
		if value == 0 {
			buffer.WriteString("_")
		} else {
			buffer.WriteString(strconv.Itoa(value))
		}
	}
	return buffer.String()
}

func NewSquare(value int) Square {
	return Square{Candidates: [9]bool{
		true,
		true,
		true,
		true,
		true,
		true,
		true,
		true,
		true},
		Value: value}
}

func WorkingFromPuzzle(p Puzzle) Working {
	var working Working
	for i, rows := range p.Initial {
		for ii, value := range rows {
			working[i][ii] = NewSquare(value)
		}
	}
	return working
}

func PuzzleFromWorking(working Working) Puzzle {
	var puzzle Puzzle
	for i, rows := range working {
		var row Row
		for ii, square := range rows {
			row[ii] = square.Value
		}
		puzzle.Initial[i] = row
	}
	return puzzle
}
