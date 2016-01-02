package solver

// Grid is a set of 81 integers representing a Sudoku board. Values 1-9
// represent given or found values. 0s represent unknowns.
type Grid [81]int

// Solver defines the interface for a Sudoku solver
type Solver interface {
	Solve(g Grid) Grid
}

// MultiSolver defines the interface for a Sudoku solver
// that finds many solutioons to a puzzle
type MultiSolver interface {
	Solve(g Grid) []Grid
}

// NewSolver creates a new Sudoku solver with teh default solving strategy.
// Currently only the backtracking solver is able to solve more than the most
// basic puzzles and so is returned.
func NewSolver() Solver {
	return NewBacktrackingSolver()
}

func NewBacktrackingSolver() Solver {
	return backtrackingSolver{}
}

func NewLogicalSolver() Solver {
	return logicalSolver{}
}

func NewRandBacktrackingSolver() Solver {
	return randBacktrackingSolver{}
}

func NewMultiBacktrackingSolver() MultiSolver {
	return multiBacktrackingSolver{}
}
