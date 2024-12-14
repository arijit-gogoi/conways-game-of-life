package main

type Grid [][]bool

// initGrid initialises the grid.
func (g *Grid) initGrid(rows, cols int) {
	*g = make([][]bool, rows)

	for r := 0; r < rows; r++ {
		(*g)[r] = make([]bool, cols)
	}
}

// makeAlive the cell in (row, col) and make it live.
func (g Grid) makeAlive(row, col int)

// kill the cell in (row, col)
func (g Grid) kill(row, col int)

// countLiveNeighbors counts the number of neighbors
// the cell in (row, col) has.
func (g Grid) countLiveNeighbors(row, col int) int

// nextGeneration is the next iteration of the game
// based on the rules of the Game.
func (g Grid) nextGeneration()
