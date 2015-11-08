package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/chrismar035/solver"
	"github.com/chrismar035/solver/backtracker"
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

		var puzzle [81]int
		for i, char := range strings.Split(line, "") {
			puzzle[i], err = strconv.Atoi(char)
			if err != nil {
				fmt.Println("Invalid puzzle line", lineNumber, "char", i)
			}
		}

		finished := solver.Puzzle{Initial: puzzle, Solution: backtracker.Solve(puzzle)}
		fmt.Println(finished)
		finished = solver.NewPuzzle(puzzle)
		fmt.Println(finished)
	}

	// var candidate [81]int
	// for i, _ := range candidate {
	// 	candidate[i] = rand.Intn(10)
	// }
	// finished := solver.NewPuzzle(candidate)
	// fmt.Println(finished)
	// fmt.Println(finished.Solved())
}
