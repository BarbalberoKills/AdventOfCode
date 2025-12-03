package p1

import (
	"fmt"
	"strconv"

	"github.com/BarbalberoKills/AdventOfCode/2025/utils"
)

const offset = 50

func Solve(input string) int {
	scanner, document, err := utils.ReadFileLighter(input)
	if err != nil {
		fmt.Println("error opening file: ", err)
	}
	defer document.Close()

	var password int
	position := offset
	for scanner.Scan() {
		line := scanner.Text()
		position = dialMove(line, position)
		// I simply check if the new position is 0, if it is I count it as valid click
		if checkDial(0, position) {
			password += 1
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("\nError reading file: %s", err)
	}
	return password
}

// for each move I read the direction and I it is positive I add it, if it's negative I subtract it
// if after the subtraction the new value is < 0 means I need to subtract 100 to come back into positive values
func dialMove(move string, position int) int {
	dir := move[:1]
	steps, _ := strconv.Atoi(move[1:])
	if dir == "R" {
		position = (position + steps) % 100
	} else {
		position = (position - steps) % 100
	}
	if position < 0 {
		position = 100 + position
	}
	return position
}

func checkDial(reference, position int) bool {
	if position == reference {
		return true
	}
	return false
}
