package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("/home/andrea/projects/AdventOfCode/2025/2/02-A_inputlist.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
	}

	var invalid int
	ranges := strings.Split(string(file), ",")
	for _, interval := range ranges {
		invalid = invalid + checkInvalid(interval)
	}
	fmt.Println(invalid)

	// fmt.Printf(strings.Join(ranges, "\n"))
	// fmt.Printf("%v    Rotatate: %v Position: %v and points at 0: %v\n\n", lineCounter, line, position, counter)

	// fmt.Println("PasswordA: ", passwordA)
	// fmt.Println("PasswordB: ", passwordB)
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
