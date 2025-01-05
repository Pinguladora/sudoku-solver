package grid

import "fmt"

type Grid [9][9]int

// Copy creates a deep copy of the grid.
func (g Grid) Copy() Grid {
	var dest Grid
	for i := range g {
		copy(dest[i][:], g[i][:])
	}
	return dest
}

// Displays the Sudoku grid.
func (g Grid) Print() {
	for _, row := range g {
		fmt.Println(row)
	}
}

var example = Grid{
	{5, 3, 0, 0, 7, 0, 0, 0, 0},
	{6, 0, 0, 1, 9, 5, 0, 0, 0},
	{0, 9, 8, 0, 0, 0, 0, 6, 0},
	{8, 0, 0, 0, 6, 0, 0, 0, 3},
	{4, 0, 0, 8, 0, 3, 0, 0, 1},
	{7, 0, 0, 0, 2, 0, 0, 0, 6},
	{0, 6, 0, 0, 0, 0, 2, 8, 0},
	{0, 0, 0, 4, 1, 9, 0, 0, 5},
	{0, 0, 0, 0, 8, 0, 0, 7, 9},
}

// Generates a dummy grid.
func ExampleGrid() Grid {
	return example
}
