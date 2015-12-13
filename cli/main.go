package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/chrismar035/solver"
)

func main() {
	file, err := os.Open("puzzles.txt")
	if err != nil {
		fmt.Println("Error opening puzzles", err)
	}
	scanner := bufio.NewScanner(file)
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		if len(line) != 81 {
			fmt.Println("Puzzle improperly sized on line", line, err)
		}

		var grid solver.Grid
		for i, char := range strings.Split(line, "") {
			grid[i], err = strconv.Atoi(char)
			if err != nil {
				fmt.Println("Invalid puzzle line", lineNumber, "char", i)
			}
		}

		puzzle := solver.Puzzle{Initial: grid}
		solver := solver.NewSolver()
		puzzle.Solution = solver.Solve(grid)
		fmt.Println(puzzle)
	}
}
