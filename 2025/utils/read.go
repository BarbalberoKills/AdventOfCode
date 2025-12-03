package utils

import (
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
