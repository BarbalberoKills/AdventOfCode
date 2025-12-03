package p1

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/BarbalberoKills/AdventOfCode/2025/utils"
)

func Solve(input string) int {
	file, err := utils.ReadFile(input)
	if err != nil {
		fmt.Println("error opening file: ", err)
	}

	var invalid int
	ranges := strings.Split(string(file), ",")
	for _, interval := range ranges {
		invalid = invalid + checkInvalid(interval)
	}
	return invalid
}

func checkInvalid(interval string) int {
	bounds := strings.Split(interval, "-")
	start, _ := strconv.Atoi(bounds[0])
	end, _ := strconv.Atoi(bounds[1])
	var invalid int

	// I cycle on the given range and for each value I splitted it in half
	// and check if the first hals is equal to the second half
	// if yes that is an invalid IDs
	for id := start; id <= end; id++ {
		idString := strconv.Itoa(id)
		if (len(idString) % 2) == 0 {
			if idString[:(len(idString)/2)] == idString[(len(idString)/2):] {
				invalid += id
			}
		}
	}
	return invalid
}
