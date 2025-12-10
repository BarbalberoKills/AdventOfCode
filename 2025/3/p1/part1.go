package p1

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/BarbalberoKills/AdventOfCode/2025/utils"
)

type Node struct {
	value    int
	position int
}

func Solve(input string) int {
	file, err := utils.ReadFile(input)
	if err != nil {
		fmt.Println("error opening file: ", err)
	}

	var totalJoltage int
	banks := strings.Split(string(file), "\n")
	for _, bank := range banks {
		totalJoltage += batteryJoltage(bank)
	}
	return totalJoltage
}

func batteryJoltage(bank string) int {

	higher := Node{value: 0, position: 0}
	secondHigher := Node{value: 0, position: 0}

	// store the higher value of the line excluding the last one
	// because I need to parse a second value
	for i, battery := range bank {
		batteryInt, _ := strconv.Atoi(string(battery))

		if batteryInt > higher.value && i != len(bank)-1 {
			higher.value = batteryInt
			higher.position = i
		}
	}

	// I scan for the second highest value in the bank starting
	// after the position of the fist value found
	for i := higher.position + 1; i < len(bank); i++ {
		batteryInt, _ := strconv.Atoi(string(bank[i]))

		if batteryInt > secondHigher.value {
			secondHigher.value = batteryInt
			secondHigher.position = i
		}
	}
	// conver both value to elements of a string
	s := []string{strconv.Itoa(higher.value), strconv.Itoa(secondHigher.value)}
	// combining them into one string only and cast it to int
	result, _ := strconv.Atoi(strings.Join(s, ""))
	return result
}
