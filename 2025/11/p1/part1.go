package p1

import (
	"bufio"
	"fmt"
	"slices"
	"strings"

	"github.com/BarbalberoKills/AdventOfCode/2025/utils"
)

type Device struct {
	name        string
	connections []*Device
}

func Solve(input string) int {
	scanner, file, err := utils.ReadFileLighter(input)
	if err != nil {
		fmt.Println("error opening file: ", err)
	}
	defer file.Close()
	devices := parseFile(scanner)
	return pathsCount(devices, "you", "out")
}

func parseFile(scanner *bufio.Scanner) []*Device {
	var devices []*Device
	for scanner.Scan() {
		line := scanner.Text()
		name := strings.Split(line, ":")
		device := &Device{name: name[0]}

		index := slices.IndexFunc(devices, func(d *Device) bool {
			return d.name == name[0]
		})
		if index == -1 {
			devices = append(devices, device)
		} else {
			device = devices[index]
		}

		connections := strings.Split(strings.TrimSpace(name[1]), " ")
		for _, c := range connections {
			index = slices.IndexFunc(devices, func(d *Device) bool {
				return d.name == c
			})
			if index != -1 {
				device.connections = append(device.connections, devices[index])
			} else {
				d := &Device{name: c}
				device.connections = append(device.connections, d)
				devices = append(devices, d)
			}
		}
	}

	// for i, d := range devices {
	// 	fmt.Printf("Name: %v - pointer: %p\n", d.name, devices[i])
	// 	for _, c := range d.connections {
	// 		fmt.Printf("	%c Child Name: %v - pointer: %p\n", '\u21b3', c.name, c)
	// 	}
	// 	fmt.Println()
	// }
	return devices
}

func ordered(devices []*Device, first string) []*Device {
	for i, d := range devices {
		if d.name == first {
			oldFirst := devices[0]
			devices[0] = d
			devices[i] = oldFirst
		}
	}
	return devices
}

func pathsCount(devices []*Device, in, out string) int {
	devices = ordered(devices, in)
	counter := checkConnection(devices[0], out)
	return counter
}

func checkConnection(device *Device, match string) int {
	counter := 0
	for _, conn := range device.connections {
		if conn.name == match {
			counter++
			return counter
		} else {
			counter += checkConnection(conn, match)
		}
	}
	return counter
}
