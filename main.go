package main

import "fmt"

type Grid [][]bool

const ROWS int = 3
const COLS int = 3
const MAX_GENERATION int = 10

var grid Grid
var tempGrid Grid

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

// liveNeighbors counts the number of neighbors
// the cell in (row, col) has.
func (g Grid) liveNeighbors(row, col int) int {
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

// Copy copies into a new grid the contents of the source.
// Used to generate next generation.
func Copy(target [][]bool, source [][]bool) {
	for r := 0; r < ROWS; r++ {
		for c := 0; c < COLS; c++ {
			target[r][c] = source[r][c]
		}
	}
}

// nextGeneration is the next iteration of the game
// based on the rules of the Game.
func (g Grid) nextGeneration() {
	Copy(tempGrid, g)

	for r := 0; r < ROWS; r++ {
		for c := 0; c < COLS; c++ {
			count := g.liveNeighbors(r, c)
			// Rule 1 and 2
			if g[r][c] == true && (count < 2 || count >= 4) {
				tempGrid[r][c] = false
			}
			// Rule 4
			if g[r][c] == false && count == 3 {
				tempGrid[r][c] = true
			}

		}
	}
	Copy(g, tempGrid)
}

func consoleLog() {
	for r := 0; r < ROWS; r++ {
		for c := 0; c < COLS; c++ {
			if grid[r][c] == true {
				fmt.Print("o ")
			} else {
				fmt.Print("- ")
			}
		}
		fmt.Println()
	}
	fmt.Println("============")
}

func main() {
	grid.initGrid(3, 3)
	tempGrid.initGrid(3, 3)

	grid.makeAlive(0, 0)
	grid.makeAlive(0, 2)
	grid.makeAlive(1, 0)
	grid.makeAlive(1, 1)
	grid.makeAlive(2, 2)
	consoleLog()

	for generation := 0; generation < MAX_GENERATION; generation++ {
		grid.nextGeneration()
		consoleLog()
	}
}
