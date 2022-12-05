package advent2022

import (
	"bufio"
	"os"
)

func GetInput(filename string) ([]string, error) {

	f, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}
