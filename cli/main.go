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
	line := 0
	for scanner.Scan() {
		line++
		line := scanner.Text()
		if len(line) != 81 {
			fmt.Println("Puzzle improperly sized on line", line, err)
		}

		var puzzle [81]int
		for i, char := range strings.Split(line, "") {
			puzzle[i], err = strconv.Atoi(char)
			if err != nil {
				fmt.Println("Invalid puzzle line", line, "char", i)
			}
		}
		finished := solver.NewPuzzle(puzzle)
		fmt.Println(finished)
		fmt.Println(finished.Solved())
	}

	// var candidate [81]int
	// for i, _ := range candidate {
	// 	candidate[i] = rand.Intn(10)
	// }
	// finished := solver.NewPuzzle(candidate)
	// fmt.Println(finished)
	// fmt.Println(finished.Solved())
}
