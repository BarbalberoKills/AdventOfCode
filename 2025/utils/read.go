package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(input string) ([]byte, error) {
	file, err := os.ReadFile(input)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	return file, nil
}

func ReadFileLighter(input string) (*bufio.Scanner, *os.File, error) {
	file, err := os.Open(input)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening file: %v", err)
	}

	scanner := bufio.NewScanner(file)

	return scanner, file, nil
}
