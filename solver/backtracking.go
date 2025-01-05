package solver

import "sudoku-solver/grid"

// Implements the basic backtracking algorithm.
type BacktrackingSolver struct{}

// Solves using basic backtracking.
func (b BacktrackingSolver) Solve(board *grid.Grid) bool {
	return backtrack(board)
}

// Raw backtracking solution.
func backtrack(board *grid.Grid) bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			// Skip already filled cells
			if board[row][col] != 0 {
				continue
			}
			return canPlaceNumber(board, row, col)
		}
	}
	return true
}

// Attempts to fill a cell in the grid.
func canPlaceNumber(board *grid.Grid, row, col int) bool {
	for num := 1; num <= 9; num++ {
		if isValidNumber(board, row, col, num) {
			board[row][col] = num
			if backtrack(board) {
				return true
			}
			// Undo movement (backtrack)
			board[row][col] = 0
		}
	}
	return false
}

// Checks if placing a number is valid.
func isValidNumber(board *grid.Grid, row, col, num int) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == num || board[i][col] == num ||
			board[3*(row/3)+i/3][3*(col/3)+i%3] == num {
			return false
		}
	}
	return true
}
