package p2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/BarbalberoKills/AdventOfCode/2025/utils"
)

type Node struct {
	joltages  [12]int
	positions [12]int
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
	// I've created a node that will stores all the number in an array
	// and their position in the bank
	batteryOn := Node{}
	for i, battery := range bank {
		batteryInt, _ := strconv.Atoi(string(battery))
		// I check each value in the line and store the HIGHEST
		// but I search only in the first 3 characters,
		// because later I need to check 11 element more
		if batteryInt > batteryOn.joltages[0] && i <= len(bank)-12 {
			batteryOn.joltages[0] = batteryInt
			batteryOn.positions[0] = i
		}
	}

	batteryOn.joltageComparetor(0, bank)

	var s []string
	for _, j := range batteryOn.joltages {
		s = append(s, strconv.Itoa(j))
	}
	result, _ := strconv.Atoi(strings.Join(s, ""))
	return result
}

// this function recursivelly search for the highest value present,
// starting only from the point where the previous highest value was
func (n *Node) joltageComparetor(counter int, bank string) {
	current := counter + 1
	if counter < 11 {
		for i := n.positions[counter] + 1; i < len(bank); i++ {
			batteryInt, _ := strconv.Atoi(string(bank[i]))

			if batteryInt > n.joltages[current] && i <= len(bank)-(12-current) {
				n.joltages[current] = batteryInt
				n.positions[current] = i
			}
		}
		n.joltageComparetor(counter+1, bank)
	}

}
