# Sudoku Solver

This generates and solved Sudoku puzzles with the goal to "discover" all 
possible puzzles and solutions. Sudoke Solver employes different algorithms
to generate and solve puzzles as described below.

## Solvers

Sudoku Solver has two classes of solvers: single and multi-solvers. Single
solvers will find a single solution for a puzzle while multi-solvers will
find multiple solutions for a single problem. The strategy of the mult-solver
will determine if it can find all solutions to a puzzle or not.

## Single Solvers

### Backtracking

This is the most naive and simplest solver. It will try every possible 
combination of values for unknown squares in order until a solution is found.
This solver is also the default solver.

```
solver := solver.NewSolver()
solution := solver.Solve(grid)

// or
solver := solver.NewBacktrackingSolver()
solution := solver.Solve(grid)
```

### Logical

The logical solver attempts to use only logic rules and cancellation to solve
puzzles. In its current state, it can only solve very easy puzzles.

```
solver := solver.NewLogicalSolver()
solution := solver.Solve(grid)
```

### Random backtracking

The random backtracking solver is similar to the backtracking solver except that
it will try values for each unknown square in random (not numerical) order. The
digits 1-9 are shuffled for each square before the solver begins and as the
solver progresses that order of digits is used.

```
solver := solver.NewRandBacktrackingSolver()
solution := solver.Solve(grid)
```

## Multi-solvers

### Backtracking

The multi-backtracking solver uses the same underlying algorithm as the single
backtracking solver except that it does not stop at the first solution. Instead,
it continues and finally returns a slice of solutions.

```
solver := solver.NewMultiBacktrackingSolver()
solutions := solver.Solve(grid)
```

## CLI Usage

Currenly, there is a very naive solver implemented. It can solve "easy" level
puzzles. To exercise it on the command line, run or build `cli/main.go`. It will
read puzzles from a `puzzles.txt` in the same directory. `puzzles.txt` should
have puzzles separated by a new line and consist of 81 consecutive digits with 0
representing blanks and 1-9 representing given squares. For example, a
`puzzles.txt` containing two puzzles (one that we can solve and one we can't):

```
010006527780145009000020010005000746000907000671000900030090000900483065168500090
750000020100200000300090406000170000001030500000048000809050002000007003060000051
```

The cli program will nicely print out the puzzle and then the solution (or as
much as could be solved) followed by true or false whether the puzzle was fully
solved. For example, the `puzzles.txt` above would produce the following output:

```
Puzzle
_1_ __6 527
78_ 145 __9
___ _2_ _1_

__5 ___ 746
___ 9_7 ___
671 ___ 9__

_3_ _9_ ___
9__ 483 _65
168 5__ _9_

Solution
419 836 527
782 145 639
356 729 418

295 318 746
843 967 251
671 254 983

534 691 872
927 483 165
168 572 394
true
Puzzle
75_ ___ _2_
1__ 2__ ___
3__ _9_ 4_6

___ 17_ ___
__1 _3_ 5__
___ _48 ___

8_9 _5_ __2
___ __7 __3
_6_ ___ _51

Solution
75_ ___ _2_
1__ 2__ ___
3__ _9_ 4_6

___ 17_ ___
__1 _3_ 5__
___ _48 ___

8_9 _5_ __2
___ __7 __3
_6_ ___ _51
false
```

## Data format

Puzzles and solutions are represented as 81 element integer arrays. With each
element, corresponding to a cell in the puzzle grid; left to right, top to
bottom.

```
 0  1  2   3  4  5   6  7  8
 9 10 11  12 13 14  15 16 17
18 19 20  21 22 23  24 25 26

27 28 29  30 31 32  33 34 35
36 37 38  39 40 41  42 43 44
45 46 47  48 49 50  51 52 53

54 55 56  57 58 59  60 61 62
63 64 65  66 67 68  69 70 71
72 73 74  75 76 77  78 79 80
```

