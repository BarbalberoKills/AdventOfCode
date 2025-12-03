package p2

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

	position := offset
	var passwordB, counter int
	lineCounter := 0

	for scanner.Scan() {
		line := scanner.Text()
		lineCounter += 1
		position, counter = dialMove(line, position)
		// fmt.Printf("%v    Rotatate: %v Position: %v and points at 0: %v\n\n", lineCounter, line, position, counter)
		// I sum all accumulated click for each movement into a total counter
		passwordB = passwordB + counter
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("\nError reading file: %s", err)
	}
	return passwordB
}

// for each move I read the direction and I it is negative I invert the value
// the key concept here is to count how many times the pointer pass to 0
// using the result from dividing the position with 100
// I condensed it into 5 cases.

// C < = -1 and starting position > 0: means we rotate it to the left and we were on positive number
// so we need to use all the rotation coming from the division, plus one more
// because we were coming from the positive and we went into the negative passing from 0

// C <= -1: starting position was exactly at 0 so we just need the coming rotation

// -1 < C <= 0 and starting position !=0: means we are rotating Left coming from positive,
// but we did less then 100 click, so we just need to count 1 passage to 0

// 0 < C < 1: means we move it between 0 and 100 still on the Right, so we don't need it

// C => 1: keeps all the possible time we turn the deal on the Right

func dialMove(move string, position int) (int, int) {
	var counter int
	dir := move[:1]
	steps, _ := strconv.Atoi(move[1:])
	if dir == "L" {
		steps = -steps
	}

	a := float64(position + steps)
	c := a / 100

	if c <= -1 && position > 0 {
		counter = int(-c) + 1
	} else if c <= -1 {
		counter = int(-c)
	} else if -1 < c && c <= 0 && position != 0 {
		counter = 1
		// The case between 0 and 1 us useless. Add just for visual completion
	} else if 0 < c && c < 1 {
		counter = 0
	} else if c >= 1 {
		counter = int(c)
	}

	// This is used just to count the new postion. Not used to count the passing to 0
	position = (position + steps) % 100
	if position < 0 {
		position = 100 + position
	}
	return position, counter
}
