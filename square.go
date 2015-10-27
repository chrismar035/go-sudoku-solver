package solver

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

func (square *square) checkIndices(indices [8]int, working working) {
	for _, index := range indices {
		neighbor := working[index]
		if neighbor.Value != 0 {
			square.Candidates[neighbor.Value-1] = false
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
