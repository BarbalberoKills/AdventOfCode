package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const offset = 50

func main() {
	document, err := os.Open("/home/andrea/projects/AdventOfCode/2025/1/01-A_inputlist.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
	}
	defer document.Close()

	position := offset
	var passwordA, passwordB, counter int
	lineCounter := 0
	scanner := bufio.NewScanner(document)
	for scanner.Scan() {
		line := scanner.Text()
		lineCounter += 1
		position, counter = dialMove(line, position)
		fmt.Printf("%v    Rotatate: %v Position: %v and points at 0: %v\n\n", lineCounter, line, position, counter)
		if position == 0 {
			passwordA += 1
		}
		passwordB = passwordB + counter
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("\nError reading file: %s", err)
	}
	fmt.Println("PasswordA: ", passwordA)
	fmt.Println("PasswordB: ", passwordB)
}

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
	} else if 0 < c && c < 1 {
		counter = 0
	} else if c >= 1 {
		counter = int(c)
	}

	position = (position + steps) % 100
	if position < 0 {
		position = 100 + position
	}
	return position, counter
}

func dialMoveCount(move string, position int) int {
	var counter int
	dir := move[:1]
	steps, _ := strconv.Atoi(move[1:])
	if dir == "L" {
		steps = -steps
	}
	if steps == 0 {
		counter = 0
		return counter
	}
	if steps < 1 {
		for tick := 0; tick > steps; tick-- {
			position -= 1
			if position == 0 || position == -100 {
				counter += 1
				position = 0
			}
		}
	} else if steps > 1 {
		for tick := 0; tick < steps; tick++ {
			position += 1
			if position == 0 || position == 100 {
				counter += 1
				position = 0
			}
		}

	}
	return counter

}

func dialMovePosition(move string, position int) int {
	dir := move[:1]
	steps, _ := strconv.Atoi(move[1:])
	if dir == "L" {
		steps = -steps
	}

	position = (position + steps) % 100
	if position < 0 {
		position = 100 + position
	}
	return position
}
