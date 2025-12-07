package p2

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

	count := navigate(matrix)

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

func navigate(matrix [][]string) int {
	// pc, _, _, ok := runtime.Caller(1)
	// details := runtime.FuncForPC(pc)
	// fmt.Printf("called from %s\n", details.Name())
	count := 0
	for r := range matrix {
		for c := range matrix[r] {
			element := matrix[r][c]

			if element == "S" {
				matrix[r+1][c] = "|"
			}

			if r < len(matrix)-1 {
				// aboveElement := matrix[r-1][c]
				beloveElement := matrix[r+1][c]
				if element == "|" && beloveElement == "." {
					matrix[r+1][c] = "|"
				} else if element == "|" && beloveElement == "^" {
					// matrix[r+1][c+1] = "|"
					matrix[r+1][c-1] = "|"
					subMatrix := matrix[r+1:]
					count += navigate(subMatrix)
					// matrix[r+1][c+1] = "|"
					// subMatrix = matrix[r+1:]
					// count += navigate(subMatrix)
				}
			}
			if r == len(matrix)-1 && element == "|" {
				count++
			}
		}
	}
	fmt.Printf("ITERATION: , COUNT:%v\n", count)
	print(matrix)
	fmt.Println()
	return count
}

func print(matrix [][]string) {
	for r := range matrix {
		fmt.Println(matrix[r])
	}
}
