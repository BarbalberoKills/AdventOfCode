package p1

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solve(input string) int {
	file, err := os.ReadFile(input)
	if err != nil {
		fmt.Println("Error opening file: ", err)
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

	for id := start; id <= end; id++ {
		idString := strconv.Itoa(id)
		if (len(idString) % 2) == 0 {
			// fmt.Printf("%v - %v - %v\n", id, len(strconv.Itoa(id)), (len(strconv.Itoa(id)) % 2))
			if idString[:(len(idString)/2)] == idString[(len(idString)/2):] {
				invalid += id
			}
		}
	}
	return invalid
}
