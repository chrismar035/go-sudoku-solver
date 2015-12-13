package solver

import (
	"bytes"
	"strconv"
)

type Puzzle struct {
	Initial  Grid
	Solution Grid
}

func (p Puzzle) IsSolved() bool {
	for _, value := range p.Solution {
		if value == 0 {
			return false
		}
	}
	return true
}

func (p Puzzle) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("Puzzle\n")
	buffer.WriteString(printInts(p.Initial))
	buffer.WriteString("\n\nSolution\n")
	buffer.WriteString(printInts(p.Solution))
	return buffer.String()
}

func printInts(ints Grid) string {
	var buffer bytes.Buffer
	for i, value := range ints {
		if i != 0 && i%9 == 0 {
			buffer.WriteString("\n")
		} else if i != 0 && i%3 == 0 {
			buffer.WriteString(" ")
		}
		if i != 0 && i%27 == 0 {
			buffer.WriteString("\n")
		}
		if value == 0 {
			buffer.WriteString("_")
		} else {
			buffer.WriteString(strconv.Itoa(value))
		}
	}
	return buffer.String()
}
