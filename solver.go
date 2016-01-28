package solver

// Grid is a set of 81 integers representing a Sudoku board. Values 1-9
// represent given or found values. 0s represent unknowns.
type Grid [81]int

// Solver defines the interface for a Sudoku solver
type Solver interface {
	Solve(g Grid) (Grid, error)
}

// MultiSolver defines the interface for a Sudoku solver
// that finds many solutioons to a puzzle
type MultiSolver interface {
	Solve(g Grid) ([]Grid, error)
}

// NewSolver creates a new Sudoku solver with the default solving strategy.
// Currently only the backtracking solver is able to solve more than the most
// basic puzzles and so is returned.
func NewSolver() Solver {
	return NewBacktrackingSolver()
}

// NewBacktrackingSolver creates a new Sudoku solver that uses a naive
// backtracking algorithm It will try each number in order until the solution
// is found or there are no options left. It mostly requires a valid puzzle.
func NewBacktrackingSolver() Solver {
	return backtrackingSolver{}
}

// NewLogicalSolver creates a new Sudoku solver that uses a logical algorithm
// to solve the puzzle. It eliminates candidates from a square based on it's
// relevant neighbors. This solver can only solve very easy puzzles.
func NewLogicalSolver() Solver {
	return logicalSolver{}
}

// NewRandBacktrackingSolver creates a new Sudoku solver that follows the naive
// backtracking algorithm, but tries to fill the squares in a random order with
// digits in random order. Providing this solver with an empty puzzle will
// return a random valid solution.
func NewRandBacktrackingSolver() Solver {
	return randBacktrackingSolver{}
}

// NewMultiBacktrackingSolver create a new Sudoku solver that repeatedly follows
// the naive backtracking algorithm until all (or no) solutions to the puzzle
// are found.
func NewMultiBacktrackingSolver() MultiSolver {
	return multiBacktrackingSolver{}
}
