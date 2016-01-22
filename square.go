package solver

import "fmt"

type square struct {
	Candidates [9]bool
	Value      int
}

func newSquare(value int) square {
	return square{Candidates: [9]bool{
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

func (square *square) checkValues(indices [8]int, working working) {
	if square.Value != 0 {
		return
	}

	for _, index := range indices {
		neighbor := working[index]
		if neighbor.Value != 0 {
			square.Candidates[neighbor.Value-1] = false
		}
	}
	square.tryToSetValueFromCandidates()
}

func (square *square) checkOtherCandidate(indices [8]int, working working) {
	if square.Value != 0 {
		return
	}

	fmt.Println(square)
	for i, possible := range square.Candidates {
		if !possible {
			continue
		}
		fmt.Println("Checking", i+1)
		found := false
		for _, index := range indices {
			neighbor := working[index]
			if neighbor.Value != 0 {
				continue
			}
			fmt.Println("Compare", index, neighbor.Candidates[i], neighbor)
			if neighbor.Candidates[i] {
				found = true
				break
			}
		}
		if !found {
			square.Value = i + 1
			return
		}
	}
}

func (square *square) tryToSetValueFromCandidates() {
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
}
