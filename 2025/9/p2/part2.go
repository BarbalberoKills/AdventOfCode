package p2

import (
	"bufio"
	"cmp"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/BarbalberoKills/AdventOfCode/2025/utils"
)

type Tile struct {
	x, y int
}

func Solve(input string) int {
	scanner, file, err := utils.ReadFileLighter(input)
	if err != nil {
		fmt.Println("error opening file: ", err)
	}
	defer file.Close()

	tiles := parseFile(scanner)

	// sort the tiles based on the x value, if x are the same, it sorts based on y
	slices.SortFunc(tiles, func(a, b Tile) int {
		c := cmp.Compare(a.x, b.x)
		if c == 0 {
			c = cmp.Compare(a.y, b.y)
		}
		return c
	})

	startRow, startCol := getStartingPoints(tiles)
	width, height := getSize(tiles)
	fmt.Printf("Starting row: %v and ending at row %v\n", startRow, height)
	fmt.Printf("Starting col: %v and ending at col %v\n", startCol, width)
	var greenTilesHorizontal []Tile
	for row := startRow; row <= height; row++ {
		// fmt.Println("Line ", row)
		check := false
		for col := startCol; col <= width; col++ {
			// fmt.Println("Column ", col)
			for i := range tiles {
				if tiles[i].x > row {
					break
				}

				// fmt.Println(tiles[i].x, tiles[i].y, check)
				if tiles[i].x == row && tiles[i].y == col && !check {
					check = true
				} else if tiles[i].x == row && tiles[i].y == col && check {
					check = false
				}
				if check {
					greenTilesHorizontal = append(greenTilesHorizontal, Tile{x: row, y: col})
				}
			}
		}
	}

	fmt.Println(greenTilesHorizontal)

	// Calculate biggest area
	biggestArea := 0.0
	for i := range tiles {
		for y := (i + 1); y <= len(tiles)-1; y++ {
			area := calcArea(tiles[i].x, tiles[i].y, tiles[y].x, tiles[y].y)
			// fmt.Printf("Start: %v End: %v Area: %v\n", tiles[i], tiles[y], area)
			if area >= biggestArea {
				biggestArea = area
			}
		}
	}

	// visualizer(greenTilesHorizontal)

	return int(biggestArea)
}

func parseFile(scanner *bufio.Scanner) []Tile {
	var tiles []Tile
	for scanner.Scan() {
		line := scanner.Text()
		posString := strings.Split(line, ",")
		var pos []int
		for _, posS := range posString {
			posInt, _ := strconv.Atoi(posS)
			pos = append(pos, posInt)
		}
		tiles = append(tiles, Tile{x: pos[1], y: pos[0]})
	}
	return tiles
}

func calcArea(x1, y1, x2, y2 int) float64 {

	area := (math.Abs((float64(x2) - float64(x1))) + 1) * (math.Abs((float64(y2) - float64(y1))) + 1)
	return area
}

func getStartingPoints(tiles []Tile) (int, int) {
	startRow := tiles[0].x
	startCol := tiles[0].y
	for i := range tiles {
		if tiles[i].x <= startRow {
			startRow = tiles[i].x
		}
		if tiles[i].y <= startCol {
			startCol = tiles[i].y
		}
		// fmt.Println(tiles[i])
	}
	return startRow, startCol
}

// Calculate max width and height
func getSize(tiles []Tile) (int, int) {
	width := 0
	height := 0
	for i := range tiles {
		if tiles[i].x >= height {
			height = tiles[i].x
		}
		if tiles[i].y >= width {
			width = tiles[i].y
		}
		// fmt.Println(tiles[i])
	}
	return width, height
}

func visualizer(tiles []Tile) {
	width, height := getSize(tiles)
	for i := range tiles {
		if tiles[i].x >= height {
			height = tiles[i].x
		}
		if tiles[i].y >= width {
			width = tiles[i].y
		}
		// fmt.Println(tiles[i])
	}

	for r := 0; r <= height; r++ {
		for c := 0; c <= width; c++ {
			// fmt.Printf("{%v %v}\n", c, r)
			// fmt.Printf("{%v %v}\n", tiles[0].x, tiles[0].y)
			check := func(tiles []Tile, col, row int) bool {
				for i := range tiles {
					if tiles[i].x == row && tiles[i].y == col {
						return true
					}
				}
				return false
			}(tiles, c, r)
			if check {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
}
