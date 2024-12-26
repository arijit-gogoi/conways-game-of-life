package main

import (
	"fmt"
	"time"
)

const ROWS int = 10
const COLS int = 10
const MAX_GENERATION int = 10

type Grid [][]bool

// initGrid initialises the grid.
func initGrid(rows, cols int) *Grid {
	g := make(Grid, rows)

	for r := 0; r < rows; r++ {
		g[r] = make([]bool, cols)
	}
	return &g
}

// makeAlive the cell in (row, col).
func (g *Grid) makeAlive(row, col int) {
	grid := *g
	grid[row][col] = true
}

// kill the cell in (row, col).
func (g *Grid) kill(row, col int) {
	grid := *g
	grid[row][col] = false
}

// liveNeighbors counts the number of neighbors
// the cell in (row, col) has.
func (g *Grid) liveNeighbors(row, col int) int {
	grid := *g
	count := 0
	// check north west
	if row > 0 && col > 0 && grid[row-1][col-1] == true {
		count++
	}

	// check north neighbor
	if row > 0 && grid[row-1][col] == true {
		count++
	}

	// check north east
	if row > 0 && col < COLS-1 && grid[row-1][col+1] == true {
		count++
	}

	// check east
	if col < COLS-1 && grid[row][col+1] == true {
		count++
	}

	// check south east
	if row < ROWS-1 && col < COLS-1 && grid[row+1][col+1] == true {
		count++
	}

	// check south
	if row < ROWS-1 && grid[row+1][col] == true {
		count++
	}

	// check south west
	if row < ROWS-1 && col > 0 && grid[row+1][col-1] == true {
		count++
	}

	// check west
	if col > 0 && grid[row][col-1] == true {
		count++
	}

	return count
}

// copyInto copies into a new grid the contents of the source.
// Used to generate next generation.
func copyInto(target, source Grid) {
	for r := 0; r < ROWS; r++ {
		for c := 0; c < COLS; c++ {
			target[r][c] = source[r][c]
		}
	}
}

// nextGeneration is the next iteration of the game
// based on the rules of the Game.
func (grid *Grid) nextGeneration(tempGrid *Grid) {
	g := *grid
	t := *tempGrid
	copyInto(t, g)

	for r := 0; r < ROWS; r++ {
		for c := 0; c < COLS; c++ {
			count := g.liveNeighbors(r, c)
			// Rule 1 and 2
			if g[r][c] == true && (count < 2 || count >= 4) {
				t[r][c] = false
			}
			// Rule 4
			if g[r][c] == false && count == 3 {
				t[r][c] = true
			}

		}
	}
	copyInto(g, t)
}

func render(grid *Grid) {
	g := *grid
	for r := 0; r < ROWS; r++ {
		for c := 0; c < COLS; c++ {
			if g[r][c] == true {
				fmt.Print("o ")
			} else {
				fmt.Print("- ")
			}
		}
		fmt.Println()
	}
	for range COLS {
		fmt.Print("==")
	}
	fmt.Println("")
}

func main() {
	grid := initGrid(ROWS, COLS)
	tempGrid := initGrid(ROWS, COLS)

	grid.makeAlive(0, 1)
	grid.makeAlive(1, 2)
	grid.makeAlive(2, 0)
	grid.makeAlive(2, 1)
	grid.makeAlive(2, 2)

	render(grid)

	for generation := 0; generation < MAX_GENERATION; generation++ {
		time.Sleep(500 * time.Millisecond)
		grid.nextGeneration(tempGrid)
		render(grid)
	}
}
