package p1

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

	biggestArea := 0.0
	for i := range tiles {
		for y := (i + 1); y <= len(tiles)-1; y++ {
			area := calcArea(tiles[i].x, tiles[i].y, tiles[y].x, tiles[y].y)
			if area >= biggestArea {
				biggestArea = area
			}
		}
	}

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
		tiles = append(tiles, Tile{x: pos[0], y: pos[1]})
	}
	return tiles
}

func calcArea(x1, y1, x2, y2 int) float64 {

	area := (math.Abs((float64(x2) - float64(x1))) + 1) * (math.Abs((float64(y2) - float64(y1))) + 1)
	return area
}
