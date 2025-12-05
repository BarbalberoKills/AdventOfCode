package p2

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/BarbalberoKills/AdventOfCode/2025/utils"
)

type Node struct {
	start int
	end   int
}

type Nodes struct {
	nodes []Node
}

func Solve(input string) int {
	scanner, file, err := utils.ReadFileLighter(input)
	if err != nil {
		fmt.Println("error opening file: ", err)
	}
	defer file.Close()

	var availableIngredients int
	var nodes Nodes

	ranges, _ := parseFile(scanner)
	nodes.orderRanges(ranges)
	availableIngredients = nodes.freshRangeAnalyzer()

	return availableIngredients
}

func parseFile(scanner *bufio.Scanner) ([]string, []string) {
	ranges, ingredients := []string{}, []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "-") {
			ranges = append(ranges, line)
		} else if strings.Contains(line, "\n") {
			continue
		} else {
			ingredients = append(ingredients, line)
		}
	}
	return ranges, ingredients

}

func (n *Nodes) orderRanges(ranges []string) []Node {
	for _, interval := range ranges {
		startS, endS := strings.Split(interval, "-")[0], strings.Split(interval, "-")[1]
		start, _ := strconv.Atoi(startS)
		end, _ := strconv.Atoi(endS)
		n.nodes = append(n.nodes, Node{start: start, end: end})
	}

	// TO CHECK
	sort.Slice(n.nodes, func(i, j int) bool {
		return n.nodes[i].start < n.nodes[j].start
	})
	return n.nodes
}

func (n *Nodes) freshRangeAnalyzer() int {
	newRanges := Nodes{nodes: []Node{Node{start: n.nodes[0].start, end: n.nodes[0].end}}}
	for actual := range n.nodes {
		// fmt.Printf("---Start: %v - End: %v\n", n.nodes[actual].start, n.nodes[actual].end)
		previous := &newRanges.nodes[len(newRanges.nodes)-1]

		if n.nodes[actual].start <= previous.end {
			if n.nodes[actual].end > previous.end {
				previous.end = n.nodes[actual].end
			}

		} else {
			newRanges.nodes = append(newRanges.nodes, Node{start: n.nodes[actual].start, end: n.nodes[actual].end})
		}
	}
	// fmt.Println("NEW SLICE")
	// for i := range newRanges.nodes {
	// 	fmt.Printf("---Start: %v - End: %v\n", newRanges.nodes[i].start, newRanges.nodes[i].end)
	// }

	fresh := 0
	for _, interval := range newRanges.nodes {
		fresh += interval.end + 1 - interval.start
	}
	return fresh
}
