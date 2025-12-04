package p2

import (
	"bufio"
	"fmt"

	"github.com/BarbalberoKills/AdventOfCode/2025/utils"
)

type Roll struct {
	presence bool
}

type Grid struct {
	grid [][]Roll
}

func Solve(input string) int {
	scanner, file, err := utils.ReadFileLighter(input)
	if err != nil {
		fmt.Println("error opening file: ", err)
	}
	defer file.Close()

	var totalRolls int

	grid := Grid{}

	row := 0
	for scanner.Scan() {
		grid.grid = append(grid.grid, []Roll{})
		line := scanner.Text()
		grid.makerGrid(row, line)
		row++
	}

	// visualPrint(grid)

	totalRolls = grid.parseGrid()

	return totalRolls
}

func test(scan *bufio.Scanner) {
	for scan.Scan() {
		fmt.Println(scan.Text())
	}
}

func (g *Grid) makerGrid(row int, line string) {
	for _, r := range line {
		if r == '@' {
			// fmt.Printf("Row: %v - Column: %v - Value: %v\n", row+1, i+1, string(r))
			g.grid[row] = append(g.grid[row], Roll{presence: true})
		} else {
			// fmt.Printf("Row: %v - Column: %v - Value: %v\n", row+1, i+1, string(r))
			g.grid[row] = append(g.grid[row], Roll{presence: false})
		}
	}
}

func (g *Grid) parseGrid() int {
	newGrid := Grid{}
	var totalRolls, nearRolls int
	validRol := 0
	for l, line := range g.grid {
		newGrid.grid = append(newGrid.grid, []Roll{})
		nearRolls = 0
		for r, roll := range line {
			// fmt.Printf("lr: (%v,%v) presence: %v\n", l, r, roll.presence)
			if roll.presence {
				nearRolls = g.checkNeighbours(l, r)
				if nearRolls < 4 {
					validRol++
					newGrid.grid[l] = append(newGrid.grid[l], Roll{presence: false})
					// fmt.Printf("x")
					continue
				}
			}
			if roll.presence {
				// fmt.Printf("@")
				newGrid.grid[l] = append(newGrid.grid[l], Roll{presence: true})
			} else {
				// fmt.Printf(".")
				newGrid.grid[l] = append(newGrid.grid[l], Roll{presence: false})
			}
		}
		// fmt.Println()
	}

	fmt.Println(validRol)
	totalRolls += validRol
	// fmt.Println()
	// visualPrint(newGrid)
	if validRol != 0 {
		validRol = newGrid.parseGrid()
		totalRolls += validRol
	}

	return totalRolls
}

func (g *Grid) checkNeighbours(posx, posy int) int {
	nearRolls := 0
	// check all the cirlce movements aroung the subject, excluding the combination that match itself
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				continue
			}
			newX := posx + x
			newY := posy + y
			// here I make sure that values we are going to check are inside the matrix, or it will arise errors
			if newX >= 0 && newY >= 0 && newX < len(g.grid) && newY < len(g.grid[0]) {
				// fmt.Printf("xy: (%v,%v) - Value: %v\n", x, y, g.grid[newX][newY].presence)
				if g.grid[newX][newY].presence {
					nearRolls++
				}
			}
		}
	}
	return nearRolls
}

// visual function as helper to visualize a matrix
func visualPrint(grid Grid) {
	for i := range grid.grid {
		for _, c := range grid.grid[i] {
			if c.presence {
				fmt.Printf("@")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
}
