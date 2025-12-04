package p1

import (
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

	totalRolls += grid.parseGrid()

	return totalRolls
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
	var nearRolls int
	validRol := 0
	for l, line := range g.grid {
		nearRolls = 0
		for r, roll := range line {
			// fmt.Printf("lr: (%v,%v) presence: %v\n", l, r, roll.presence)
			if roll.presence {
				nearRolls = g.checkNeighbours(l, r)
				if nearRolls < 4 {
					validRol++
					fmt.Printf("x")
					continue
				}
			}
			if roll.presence {
				fmt.Printf("@")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
		// break
	}
	return validRol
}

func (g *Grid) checkNeighbours(posx, posy int) int {
	nearRolls := 0
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				continue
			}
			// fmt.Printf("xy: (%v,%v)\n", x, y)
			newX := posx + x
			newY := posy + y
			// fmt.Printf("newX: %v, newY: %v", newX, newY)
			if newX >= 0 && newY >= 0 && newX < len(g.grid) && newY < len(g.grid[0]) {
				// fmt.Printf("xy: (%v,%v) - Value: %v\n", x, y, g.grid[newX][newY].presence)
				if g.grid[newX][newY].presence {
					nearRolls++
				}
			}
		}
	}
	// fmt.Println(nearRolls)
	// fmt.Println()
	return nearRolls
}
