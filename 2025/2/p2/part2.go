package p2

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/BarbalberoKills/AdventOfCode/2025/utils"
)

func Solve(input string) int {
	file, err := utils.ReadFile(input)
	if err != nil {
		fmt.Println("error opening file: ", err)
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

	// I cycle on the given range and for each value
	// I loop over the lenght of the string and with the module
	// I search for the value that can be divide withot given rest.
	// This give me all the possible pattern lenght I should look for possible pattern repetition.

	// fmt.Printf("Start: %v - End: %v\n", start, end)
	for id := start; id <= end; id++ {
		idString := strconv.Itoa(id) //id casted to string
		lenA := len(idString)        //count total number of element in that string
		patternLenght := []int{}
		for i := 1; i < lenA; i++ {
			if lenA%i == 0 {
				patternLenght = append(patternLenght, i)
			}
		}

		// now that I know the pattern lenght I cycle over all the pattern possibility lenght for any given value
		// for each one I create a reference pattern(marker), as bigger as the current patternLenght(starting from the bigger one)
		// Now I split the value into bits sized as the pattern size, and I check that is chuck is equal to the reference pattern.
		// If even one is differnt I set checker to false and pass to the next value, that one is a valid ID.
		slices.Reverse(patternLenght)
		for _, comb := range patternLenght {
			checker := false
			marker := idString[:comb]
			for s := 0; s < lenA; s += comb {
				e := s + comb
				if idString[s:e] == marker {
					checker = true
				} else {
					checker = false
					break
				}
			}
			if checker {
				// fmt.Printf("invalid IDs: %v\n", id)
				invalid = invalid + id
				break
			}
		}
	}
	// fmt.Printf("Range sum: %v\n\n", invalid)
	return invalid
}
