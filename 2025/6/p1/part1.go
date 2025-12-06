package p1

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"

	"github.com/BarbalberoKills/AdventOfCode/2025/utils"
)

func Solve(input string) int {
	scanner, file, err := utils.ReadFileLighter(input)
	if err != nil {
		fmt.Println("error opening file: ", err)
	}
	defer file.Close()

	numbers, operations := parseFile(scanner)
	numbers = transpose(numbers)

	total := 0
	for i, o := range operations {
		parzial := 0
		if o == "+" {
			for _, n := range numbers[i] {
				num, _ := strconv.Atoi(n)
				parzial = parzial + num
			}
		} else if o == "*" {
			for _, n := range numbers[i] {
				num, _ := strconv.Atoi(n)
				if parzial == 0 {
					parzial = num
				} else {
					parzial = parzial * num
				}
			}
		}
		// fmt.Println(parzial)
		total = total + parzial
	}

	return total
}

func parseFile(scanner *bufio.Scanner) ([][]string, []string) {
	var numbers [][]string
	var operation []string
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		r1 := regexp.MustCompile(`\d+`)
		r2 := regexp.MustCompile(`[+]|[*]`)
		rowNumbers := r1.FindAllString(line, -1)
		rowOperations := r2.FindAllString(line, -1)
		if len(rowNumbers) != 0 {
			numbers = append(numbers, rowNumbers)
		} else {
			operation = rowOperations
		}
		count++
	}
	return numbers, operation
}

func transpose(matrix [][]string) [][]string {
	m := len(matrix)
	n := len(matrix[0])

	newMatrix := make([][]string, n)
	for i := 0; i < n; i++ {
		newMatrix[i] = make([]string, m)
	}

	var row []string
	for i := 0; i < n; i++ {
		row = newMatrix[i]
		for j := 0; j < m; j++ {
			row[j] = matrix[j][i]
		}
	}

	return newMatrix
}
