package solver

type Square struct {
	Candidates [9]bool
	Value      int
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
