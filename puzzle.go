package solver

import (
	"bytes"
	"strconv"
)

type Puzzle struct {
	Initial  [81]int
	Solution [81]int
}

func (p Puzzle) String() string {
	var buffer bytes.Buffer
	for i, value := range p.Initial {
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
