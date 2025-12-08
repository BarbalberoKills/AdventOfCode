package p1

import (
	"bufio"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/BarbalberoKills/AdventOfCode/2025/utils"
)

type Box struct {
	position []int
}

type Circuit struct {
	boxes []Box
}

func Solve(input string) int {
	scanner, file, err := utils.ReadFileLighter(input)
	if err != nil {
		fmt.Println("error opening file: ", err)
	}
	defer file.Close()

	boxes := parseFile(scanner)

	// create a map that has as key the distance between the two boxes stored as values
	var distances = make(map[float64][]Box)
	for i := 0; i < len(boxes); i++ {
		for y := i + 1; y < len(boxes); y++ {
			d := calcDistance(boxes[y], boxes[i])
			distances[d] = []Box{boxes[y], boxes[i]}
		}
	}

	// sort the map based on the key, so the closest is first
	keys := make([]float64, 0, len(distances))
	for k := range distances {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	// for _, i := range keys {
	// 	fmt.Printf("Key: %v - Value: %v\n", i, distances[i])
	// }

	var circuits []Circuit

	check := true
	count := 9
	for check {
		if count == 0 {
			check = false
		}

		a := distances[keys[int(math.Abs(float64(count-9)))]][0]
		b := distances[keys[int(math.Abs(float64(count-9)))]][1]
		if len(circuits) == 0 {
			circuits = append(circuits, Circuit{boxes: []Box{a, b}})
			count--
			continue
		}
		fmt.Printf("Dist: %v - a: %v - b: %v\n", keys[int(math.Abs(float64(count-9)))], a, b)
		for i := range circuits {
			fmt.Println(count)
			// fmt.Println(c.boxes[0])
			if containsBox(circuits[i].boxes, a) && containsBox(circuits[i].boxes, b) {
				fmt.Println("z")
				break
			} else if containsBox(circuits[i].boxes, a) {
				fmt.Println("a")
				circuits[i].boxes = append(circuits[i].boxes, b)
				break
			} else if containsBox(circuits[i].boxes, b) {
				fmt.Println("b")
				circuits[i].boxes = append(circuits[i].boxes, a)
				break
			} else {
				fmt.Println("c")
				circuits = append(circuits, Circuit{boxes: []Box{a, b}})
				break
			}
		}
		// fmt.Println(circuits)
		// if count == 1 {
		// 	os.Exit(1)
		// }
		// fmt.Println(distances[keys[int(math.Abs(float64(count-9)))]][0])
		count--
	}

	for _, c := range circuits {
		fmt.Println(c)
	}

	total := len(boxes)

	return total
}

func parseFile(scanner *bufio.Scanner) []Box {
	var boxes []Box
	for scanner.Scan() {
		line := scanner.Text()
		posString := strings.Split(line, ",")
		var pos []int
		for _, ns := range posString {
			n, _ := strconv.Atoi(ns)
			pos = append(pos, n)
		}
		boxes = append(boxes, Box{position: pos})
	}
	return boxes
}

func calcDistance(first, second Box) float64 {
	x := math.Pow(float64(second.position[0]-first.position[0]), 2)
	y := math.Pow(float64(second.position[1]-first.position[1]), 2)
	z := math.Pow(float64(second.position[2]-first.position[2]), 2)
	return math.Sqrt(x + y + z)
}

func boxEqual(a, b Box) bool {
	if len(a.position) != len(b.position) {
		return false
	}
	for i := range a.position {
		if a.position[i] != b.position[i] {
			return false
		}
	}
	return true
}

func containsBox(boxes []Box, target Box) bool {
	for _, b := range boxes {
		if boxEqual(b, target) {
			return true
		}
	}
	return false
}
