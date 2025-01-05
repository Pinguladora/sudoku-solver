package solver

import "sudoku-solver/grid"

// Interface for all solvers.
type Solver interface {
	Solve(*grid.Grid) bool
}
