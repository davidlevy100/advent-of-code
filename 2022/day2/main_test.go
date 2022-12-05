package main

import (
	"os"
	"path/filepath"
	"testing"

	util "github.com/davidlevy100/advent-of-code/util"
)

var Part1 = part1
var Part2 = part2

func TestPart1(t *testing.T) {

	path, _ := os.Getwd()
	fullPath := filepath.Join(path, "input_test.txt")

	data, _ := util.GetInput(fullPath)

	ans := Part1(data)
	correct := 15

	if ans != correct {
		t.Errorf("Part 2 = %d, want %d", ans, correct)
	}

}

func TestPart2(t *testing.T) {

	path, _ := os.Getwd()
	fullPath := filepath.Join(path, "input_test.txt")

	data, _ := util.GetInput(fullPath)

	ans := Part2(data)
	correct := 12

	if ans != correct {
		t.Errorf("Part 1 = %d, want %d", ans, correct)
	}

}
