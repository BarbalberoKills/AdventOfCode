package p2

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
		idString := strconv.Itoa(id) //id casted to string
		// fmt.Printf("%v - %v - %v\n", id, len(strconv.Itoa(id)), (len(strconv.Itoa(id)) % 2))
		r := []rune(idString)
		pos := strings.IndexRune(idString[1:], r[0]) + 1
		if pos != 0 {
			fmt.Printf("Element: %v - Char: %v - Pos: %v\n", id, string(r[0]), pos)
		}
		break
		// strings.
		// if idString[:(len(idString)/2)] == idString[(len(idString)/2):] {
		// 	invalid += id
		// }
	}
	return invalid
}
