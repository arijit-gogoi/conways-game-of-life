package main

type Grid [][]bool

const ROWS int = 10
const COLS int = 10

// initGrid initialises the grid.
func (g *Grid) initGrid(rows, cols int) {
	*g = make([][]bool, rows)

	for r := 0; r < rows; r++ {
		(*g)[r] = make([]bool, cols)
	}
}

// makeAlive the cell in (row, col) and make it live.
func (g Grid) makeAlive(row, col int) {
	g[row][col] = true
}

// kill the cell in (row, col)
func (g Grid) kill(row, col int) {
	g[row][col] = false
}

// countLiveNeighbors counts the number of neighbors
// the cell in (row, col) has.
func (g Grid) countLiveNeighbors(row, col int) int {
	count := 0
	// check north west
	if row > 0 && col > 0 && g[row-1][col-1] == true {
		count++
	}

	// check north neighbor
	if row > 0 && g[row-1][col] == true {
		count++
	}

	// check north east
	if row > 0 && col < COLS-1 && g[row-1][col+1] == true {
		count++
	}

	// check east
	if col < COLS-1 && g[row][col+1] == true {
		count++
	}

	// check south east
	if row < ROWS-1 && col < COLS-1 && g[row+1][col+1] == true {
		count++
	}

	// check south
	if row < ROWS-1 && g[row+1][col] == true {
		count++
	}

	// check south west
	if row < ROWS-1 && col > 0 && g[row+1][col-1] == true {
		count++
	}

	// check west
	if col > 0 && g[row][col-1] == true {
		count++
	}

	return count

}

// nextGeneration is the next iteration of the game
// based on the rules of the Game.
func (g Grid) nextGeneration()
