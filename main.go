package main

import "fmt"

func main() {
	for i := 0; i < 256; i++ {
		fmt.Printf("*** Rule %d ***\n", i)
		printECA(i)
		fmt.Println()
	}
}

// printECA prints the space-time diagram for an ECA, starting with a single
// 'on' input
func printECA(rule int) {
	const n = 20 // number of iterations

	// The output width will be 2*n - 1. Add n cells on each side to avoid edge
	// effects from affecting the diagram
	width := 2*n - 1 + 2*n

	grid := make([]cell, width)
	grid[width/2].on = true

	for i := 0; i < n; i++ {
		// print 2n-1 cells in the center
		printGrid(grid[n : 3*n-1])
		grid = nextGen(grid, rule)
	}
}

type cell struct {
	on bool
}

// nextGen returns the next generation for an ECA, for a given rule
func nextGen(grid []cell, rule int) []cell {
	next := make([]cell, len(grid))

	for i := range grid {
		var state byte
		if i > 0 && grid[i-1].on {
			state |= 0b_100
		}
		if grid[i].on {
			state |= 0b_010
		}
		if i+1 < len(grid) && grid[i+1].on {
			state |= 0b_001
		}
		// next generation is "on" if rule has the state bit on
		next[i].on = rule&(1<<state) != 0
	}
	return next
}

func printGrid(grid []cell) {
	for _, c := range grid {
		if c.on {
			fmt.Print("#")
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
