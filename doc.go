/*
Package solver provides different algorithms for solving Sudoku puzzles.  The
currently implemented algorithms are:

Basic Backtacking

This algorighm is the most basic form of backtracking algorithm possible. It
tries each possible number in order top to bottom, left to right. When an
incorrect number is attempted, the algorithm backtracks to the previous,
non-given space and attempts the next number.

Logical (Incomplete)

The logical solver attempts to solve sudoku puzzles without backtracking using
the other squares in the same row, column, and 3x3 square as each space and
eliminating conflicting values.

This solver currently only solves very easy puzzles but can utilize more
advanced logic techniques to solve harder puzzles. It is known that there are
certain puzzles which can only be solved with backtracking. We'll see...

The logical al
*/
package solver
