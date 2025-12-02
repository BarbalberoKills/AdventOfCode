package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type number struct {
	value   int
	pointer *number
}

const offset = 50

func main() {
	// var dial [100]*number
	// for i := range dial {
	// 	dial[i] = &number{value: i}
	// 	if i != 0 {
	// 		dial[i-1].pointer = dial[i]
	// 	}
	// 	if i == (len(dial) - 1) {
	// 		dial[i].pointer = dial[0]
	// 	}
	// }

	document, err := os.Open("/home/andrea/projects/AdventOfCode/2025/1/01-A_inputlist.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
	}
	defer document.Close()

	scanner := bufio.NewScanner(document)

	var password int
	position := offset
	for scanner.Scan() {
		line := scanner.Text()
		position = dialMove(line, position)
		if checkDial(0, position) {
			password += 1
		}
		// fmt.Println(position)
		// fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("\nError reading file: %s", err)
	}
	fmt.Println(password)
}

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
