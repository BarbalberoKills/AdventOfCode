package p1

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	value    int
	position int
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
	}
	return totalJoltage
}

func batteryJoltage(bank string) int {

	higher := Node{value: 0, position: 0}
	secondHigher := Node{value: 0, position: 0}
	for i, battery := range bank {
		batteryInt, _ := strconv.Atoi(string(battery))

		if batteryInt > higher.value && i != len(bank)-1 {
			higher.value = batteryInt
			higher.position = i
		}
	}
	for i := higher.position; i < len(bank); i++ {
		batteryInt, _ := strconv.Atoi(string(bank[i]))

		if batteryInt > secondHigher.value && i > higher.position {
			secondHigher.value = batteryInt
			secondHigher.position = i
		}
	}
	s := []string{strconv.Itoa(higher.value), strconv.Itoa(secondHigher.value)}
	x := strings.Join(s, "")
	result, _ := strconv.Atoi(x)
	fmt.Println(x)
	return result
}
