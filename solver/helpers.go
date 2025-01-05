package solver

// Calculates the next cell to process.
func getNextCell(row, col int) (int, int) {
	col++
	if col == 9 {
		row++
		col = 0
	}
	return row, col
}

// Calculates the index of the 3x3 sub-grid for a given cell.
func getSubGridIndex(row, col int) int {
	return (row/3)*3 + col/3
}

// Memoization related

// Checks if a number can be placed in a given cell.
func canPlaceNumber(memo Memo, row, col, num int) bool {
	subGridIndex := getSubGridIndex(row, col)
	return !memo.Rows[row][num] && !memo.Cols[col][num] && !memo.SubGrids[subGridIndex][num]
}
