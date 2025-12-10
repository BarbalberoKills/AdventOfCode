package p1

import (
	"bufio"
	"fmt"

	"github.com/BarbalberoKills/AdventOfCode/2025/utils"
)

func Solve(input string) int {
	scanner, file, err := utils.ReadFileLighter(input)
	if err != nil {
		fmt.Println("error opening file: ", err)
	}
	defer file.Close()

	matrix := parseFile(scanner)

	count := 0
	for r := range matrix {
		// fmt.Println(matrix[r])
		for c := range matrix[r] {
			element := matrix[r][c]

			if element == "S" {
				matrix[r+1][c] = "|"
				continue
			}

			if r >= 1 && r < len(matrix)-1 {
				// aboveElement := matrix[r-1][c]
				beloveElement := matrix[r+1][c]
				if element == "|" && beloveElement == "^" {
					count++
					matrix[r+1][c+1] = "|"
					matrix[r+1][c-1] = "|"
				} else if element == "|" && beloveElement == "." {
					matrix[r+1][c] = "|"
				}
			}
		}
		// fmt.Println(matrix[r])
		// fmt.Println(count)
	}

	for r := range matrix {
		fmt.Println(matrix[r])
	}

	total := count

	return total
}

func parseFile(scanner *bufio.Scanner) [][]string {
	var matrix [][]string
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]string, 0, len(line))
		for _, r := range line {
			row = append(row, string(r))
		}
		matrix = append(matrix, row)
	}
	return matrix
}
