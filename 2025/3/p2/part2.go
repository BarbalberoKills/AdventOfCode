package p2

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	joltages  [12]int
	positions [12]int
}

func Solve(input string) int {
	file, err := os.ReadFile(input)
	if err != nil {
		fmt.Println("Error opening file: ", err)
	}

	var totalJoltage int
	banks := strings.Split(string(file), "\n")
	for _, bank := range banks {
		totalJoltage += batteryJoltage(bank)
		// break
	}
	return totalJoltage
}

func batteryJoltage(bank string) int {
	batteryOn := Node{}
	for i, battery := range bank {
		batteryInt, _ := strconv.Atoi(string(battery))

		if batteryInt > batteryOn.joltages[0] && i <= len(bank)-12 {
			batteryOn.joltages[0] = batteryInt
			batteryOn.positions[0] = i
		}
	}
	batteryOn.joltageComparetor(0, bank)

	// for i := higher.position + 1; i < len(bank); i++ {
	// 	batteryInt, _ := strconv.Atoi(string(bank[i]))

	// 	if batteryInt > secondHigher.value {
	// 		secondHigher.value = batteryInt
	// 		secondHigher.position = i
	// 	}
	// }
	// s := []string{strconv.Itoa(higher.value), strconv.Itoa(secondHigher.value)}
	// x := strings.Join(s, "")
	var s []string
	for _, j := range batteryOn.joltages {
		s = append(s, strconv.Itoa(j))
	}
	result, _ := strconv.Atoi(strings.Join(s, ""))
	// fmt.Println(batteryOn.positions)
	// fmt.Println()
	return result
}

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
