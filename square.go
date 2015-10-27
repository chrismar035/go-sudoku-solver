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
