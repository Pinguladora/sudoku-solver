package main

import (
	"fmt"
	"sudoku-solver/grid"
	"sudoku-solver/solver"
)

func main() {
	sudoku := grid.ExampleGrid()

	solverInstance := solver.BacktrackingSolver{}

	fmt.Println("Given Sudoku board:")
	sudoku.Print()

	// Solve the puzzle
	if solverInstance.Solve(&sudoku) {
		fmt.Println("\nSolved Sudokuboard:")
	} else {
		fmt.Println("\nNo solution exists. Board looks as:")
	}
	sudoku.Print()
}
