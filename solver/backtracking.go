package solver

import "sudoku-solver/grid"

// Implements the basic backtracking algorithm.
type BacktrackingSolver struct{}

// Solves using basic backtracking.
func (b BacktrackingSolver) Solve(board *grid.Grid) bool {
	return backtrack(board, 0, 0)
}

// Checks if placing a number is valid.
func isValidNumber(board grid.Grid, row, col, num int) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == num || board[i][col] == num ||
			board[3*(row/3)+i/3][3*(col/3)+i%3] == num {
			return false
		}
	}
	return true
}

// Perform backtrack in a recursive approach
func backtrack(board *grid.Grid, row, col int) bool {
	// Base case
	if row == 9 {
		return true
	}

	nextRow, nextCol := getNextCell(row, col)

	if board[row][col] != 0 {
		return backtrack(board, nextRow, nextCol)
	}
	for num := 1; num <= 9; num++ {
		if isValidNumber(*board, row, col, num) {
			board[row][col] = num
			if backtrack(board, nextRow, nextCol) {
				return true
			}
			board[row][col] = 0
		}
	}
	return false
}

// Implements the memoized backtracking algorithm.
type MemoizedBacktrackingSolver struct{}

// Solves the Sudoku puzzle using backtracking with memoization.
func (m MemoizedBacktrackingSolver) Solve(grid *grid.Grid) bool {
	memo := NewMemo(grid)
	return backtrackWithMemo(grid, &memo, 0, 0)
}

// Memo holds the memoization maps for rows, columns, and sub-grids.
type Memo struct {
	Rows     [9][10]bool // Tracks numbers placed in each row
	Cols     [9][10]bool // Tracks numbers placed in each column
	SubGrids [9][10]bool // Tracks numbers placed in each 3x3 sub-grid
}

// Initializes the memoization structure from the Sudoku grid.
func NewMemo(board *grid.Grid) Memo {
	memo := Memo{}
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			num := board[row][col]
			if num != 0 {
				memo.Rows[row][num] = true
				memo.Cols[col][num] = true
				memo.SubGrids[getSubGridIndex(row, col)][num] = true
			}
		}
	}
	return memo
}

// Recursive function for memoized backtracking.
func backtrackWithMemo(board *grid.Grid, memo *Memo, row, col int) bool {
	// Base case: Completed the grid
	if row == 9 {
		return true
	}

	// Move to the next cell
	nextRow, nextCol := getNextCell(row, col)

	if board[row][col] != 0 {
		return backtrackWithMemo(board, memo, nextRow, nextCol)
	}

	// Try placing numbers 1 to 9
	for num := 1; num <= 9; num++ {
		if canPlaceNumber(*memo, row, col, num) {
			placeNumber(board, memo, row, col, num)
			if backtrackWithMemo(board, memo, nextRow, nextCol) {
				return true
			}
			removeNumber(board, memo, row, col, num)
		}
	}

	return false // Trigger backtracking
}
