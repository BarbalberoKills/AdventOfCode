package p2

import (
	"bufio"
	"fmt"
	"slices"
	"strings"
	"time"

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
	result := timeit(func() int {
		res1 := pathsCount(devices, "svr", "fft")
		res2 := pathsCount(devices, "fft", "dac")
		res3 := pathsCount(devices, "dac", "out")
		return res1 * res2 * res3
	})
	return result
}

func timeit(f func() int) int {
	start := time.Now()
	result := f()
	elapsed := time.Since(start)
	fmt.Printf("Elapsed time: %v\n", elapsed)
	return result
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
	cache := make(map[string]int)
	counter := checkConnection(devices[0], out, cache)
	return counter
}

func checkConnection(device *Device, match string, cache map[string]int) int {
	if device.name == match {
		return 1
	}
	if v, ok := cache[device.name]; ok {
		return v
	}
	counter := 0
	for _, conn := range device.connections {
		counter += checkConnection(conn, match, cache)
	}
	cache[device.name] = counter

	return counter
}
