package solver

import "sudoku-solver/grid"

// Memoization related

// Updates the grid and memoization structure to place a number.
func placeNumber(board *grid.Grid, memo *Memo, row, col, num int) {
	board[row][col] = num
	memo.Rows[row][num] = true
	memo.Cols[col][num] = true
	memo.SubGrids[getSubGridIndex(row, col)][num] = true
}

// Undoes a placement in the grid and updates memoization.
func removeNumber(board *grid.Grid, memo *Memo, row, col, num int) {
	board[row][col] = 0
	memo.Rows[row][num] = false
	memo.Cols[col][num] = false
	memo.SubGrids[getSubGridIndex(row, col)][num] = false
}
