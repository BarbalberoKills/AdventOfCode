package p2

import (
	"bufio"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/BarbalberoKills/AdventOfCode/2025/utils"
)

func Solve(input string) int {
	scanner, file, err := utils.ReadFileLighter(input)
	if err != nil {
		fmt.Println("error opening file: ", err)
	}
	defer file.Close()

	numbers, operations := parseFile(scanner)
	slices.Reverse(operations)

	totalRow := len(numbers)
	totalCol := len(numbers[0])

	table := map[string][]int{}
	count := 0

	for col := totalCol - 1; col >= 0; col-- {
		separetor := true
		var current []string
		// here a scan the matrix from top right to bottom left each column at time,
		// if at least one element is != from 0 that is a valid colum to sum/multiply
		// so I store those values in a slice current and then as value of map that has
		// as key a string with the operation to do and a counter(keys must be unique)
		for row := range totalRow {
			num := numbers[row][col]
			current = append(current, num)
			if num != " " {
				separetor = false
			}
		}

		// here I add the slice current to the map if that is a valid range,
		// or I increment the count so the operator slice that I'll use as key for the next valuse
		if !separetor {
			c := func(numbers []string) int {
				sum := ""
				for _, n := range numbers {
					if n != " " {
						sum += n
					}
				}
				sumInt, _ := strconv.Atoi(sum)
				return sumInt
			}(current)

			table[operations[count]+strconv.Itoa(count)] = append(table[operations[count]+strconv.Itoa(count)], c)
		} else {
			count++
		}
	}

	total := 0
	for k, v := range table {
		parzial := 0
		if strings.Contains(k, "+") {
			// fmt.Printf("To SUM :%v\n", v)
			for _, n := range v {
				parzial = parzial + n
			}
		} else if strings.Contains(k, "*") {
			// fmt.Printf("To Multiply :%v\n", v)
			for _, n := range v {
				if parzial == 0 {
					parzial = 1
				}
				parzial = parzial * n
			}
		}
		fmt.Println(parzial)
		total = total + parzial
	}

	return total
}

func parseFile(scanner *bufio.Scanner) ([][]string, []string) {
	var numbers [][]string
	var operation []string
	for scanner.Scan() {
		line := scanner.Text()

		r2 := regexp.MustCompile(`[+]|[*]`)
		rowOperations := r2.FindAllString(line, -1)
		if len(rowOperations) != 0 {
			operation = rowOperations
			continue
		}

		row := make([]string, 0, len(line))
		for _, r := range line {
			row = append(row, string(r))
		}
		numbers = append(numbers, row)
	}
	return numbers, operation
}
