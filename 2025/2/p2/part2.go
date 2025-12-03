package p2

import (
	"fmt"
	"os"
	"slices"
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
		// break
	}

	// a := "222222"
	// lenA := len(a)
	// plausibleLenght := []int{}
	// for i := 1; i < lenA; i++ {
	// 	if lenA%i == 0 {
	// 		plausibleLenght = append(plausibleLenght, i)
	// 		// fmt.Printf("%v has %v units and rest of division for %v is %v\n", a, lenA, i, lenA%i)
	// 	}
	// }
	// slices.Reverse(plausibleLenght)
	// for _, comb := range plausibleLenght {
	// 	fmt.Println(comb)
	// 	checker := false
	// 	marker := a[:comb]
	// 	for s := 0; s < lenA; s += comb {
	// 		e := s + comb
	// 		if a[s:e] == marker {
	// 			checker = true
	// 		} else {
	// 			checker = false
	// 			break
	// 		}
	// 		fmt.Println(checker)
	// 	}
	// 	fmt.Println(checker)
	// 	if checker {
	// 		fmt.Println(a)
	// 		break
	// 	}

	// }
	return invalid
}

func checkInvalid(interval string) int {

	bounds := strings.Split(interval, "-")
	start, _ := strconv.Atoi(bounds[0])
	end, _ := strconv.Atoi(bounds[1])
	var invalid int
	fmt.Printf("Start: %v - End: %v\n", start, end)
	for id := start; id <= end; id++ {
		// fmt.Println(id)
		idString := strconv.Itoa(id) //id casted to string
		lenA := len(idString)
		plausibleLenght := []int{}
		for i := 1; i < lenA; i++ {
			if lenA%i == 0 {
				plausibleLenght = append(plausibleLenght, i)
			}
		}

		slices.Reverse(plausibleLenght)
		for _, comb := range plausibleLenght {
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
				fmt.Printf("invalid IDs: %v\n", id)
				invalid = invalid + id
				break
			}
		}
	}
	fmt.Printf("Range sum: %v\n\n", invalid)
	return invalid
}
