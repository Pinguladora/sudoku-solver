package main

import (
	"fmt"
	"sudoku-solver/grid"
	"sudoku-solver/solver"
)

func main() {
	sudoku := grid.ExampleGrid()
	fmt.Println("Given Sudoku board:")
	sudoku.Print()

	solverInstance := solver.MemoizedBacktrackingSolver{}

	if solverInstance.Solve(&sudoku) {
		fmt.Println("\nSolved Sudoku board:")
	} else {
		fmt.Println("\nNo solution exists. Board looks as:")
	}
	sudoku.Print()
}
