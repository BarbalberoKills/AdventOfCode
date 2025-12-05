package p1

import (
	"bufio"
	"fmt"
	"slices"
	"strconv"
	"strings"

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

	var availableIngredients int

	ranges, ingredients := parseFile(scanner)
	// fmt.Println(strings.Join(ranges, ""))
	// fmt.Println("CIAO")
	// fmt.Println(strings.Join(ingredients, ""))

	availableIngredients = len(freshCheck(ranges, ingredients))

	return availableIngredients
}

func parseFile(scanner *bufio.Scanner) ([]string, []string) {
	ranges, ingredients := []string{}, []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "-") {
			ranges = append(ranges, line)
		} else if strings.Contains(line, "\n") {
			continue
		} else {
			ingredients = append(ingredients, line)
		}
	}
	return ranges, ingredients

}

func freshCheck(ranges, ingredients []string) []int {
	checkedIng := []int{}
	for _, interval := range ranges {
		fmt.Printf("Checking range: %v\n", interval)
		startS, endS := strings.Split(interval, "-")[0], strings.Split(interval, "-")[1]
		start, _ := strconv.Atoi(startS)
		end, _ := strconv.Atoi(endS)
		// fmt.Printf("Start: %v - End: %v\n", start, end)
		for _, ingredientS := range ingredients {
			ingredient, _ := strconv.Atoi(ingredientS)
			if start <= ingredient && ingredient <= end && !slices.Contains(checkedIng, ingredient) {
				checkedIng = append(checkedIng, ingredient)
			}
		}
	}
	return checkedIng
}
