package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {
	const n = 32 // number of iterations

	// Generate color and bw pngs for all elementary cellular automata
	for i := 0; i < 256; i++ {
		fmt.Printf("*** Rule %d ***\n", i)
		//printECA(i, n)
		if err := plotECA(fmt.Sprintf("rule%.3d_color.png", i), i, n, true); err != nil {
			log.Fatal(err)
		}
		if err := plotECA(fmt.Sprintf("rule%.3d_bw.png", i), i, n, false); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("done")
}

// plotECA plots the space-time diagram for an ECA, starting with a single
// 'on' input
// The result is stored as a png file
func plotECA(filename string, rule, n int, useColor bool) error {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// The output width will be n. Add n cells on each side to avoid edge
	// effects from affecting the diagram
	width := 3 * n

	grid := make([]cell, width)
	grid[width/2].on = true

	pal := color.Palette([]color.Color{
		color.RGBA{85, 85, 85, 255},   // Gray
		color.RGBA{255, 85, 85, 255},  // red
		color.RGBA{85, 255, 85, 255},  // green
		color.RGBA{255, 255, 85, 255}, // yellow
		color.RGBA{85, 85, 255, 255},  // blue
		color.RGBA{255, 85, 255, 255}, // magenta
		color.RGBA{85, 255, 255, 255}, // cyan
		color.Black,
		color.White,
	})

	size := 5 // cell size in pixels
	img := image.NewPaletted(image.Rect(0, 0, n*size, n*size), pal)
	b := img.Bounds()

	xoff := (len(grid) - n) / 2 // offset to for n cells in the center
	for y := 0; y < n; y++ {
		for x := 0; x < n; x++ {
			c := grid[x+xoff]
			for dx := 0; dx < size; dx++ {
				for dy := 0; dy < size; dy++ {
					if c.on {
						if useColor {
							img.SetColorIndex(x*size+dx, y*size+dy, c.prevgen)
						} else {
							img.Set(x*size+dx, y*size+dy, color.Black)
						}
					} else {
						img.Set(x*size+dx, y*size+dy, color.White)
					}
				}
			}
		}
		grid = nextGen(grid, rule)
	}

	if err := png.Encode(f, img.SubImage(b)); err != nil {
		return err
	}
	return nil
}

// printECA prints the space-time diagram for an ECA, starting with a single
// 'on' input
func printECA(rule, n int) {
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
	on      bool
	prevgen byte
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
		next[i].prevgen = state
	}
	return next
}

// printGrid prints an ECA grid as a line to stdout
func printGrid(grid []cell) {
	for _, c := range grid {
		if c.on {
			fmt.Printf("%d", c.prevgen)
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
