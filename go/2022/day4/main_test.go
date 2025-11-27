package main

import (
	"log"
	"testing"

	util "github.com/davidlevy100/advent-of-code/util"
)

var Part1 = part1
var Part2 = part2

func TestPart1(t *testing.T) {

	data, err := util.GetInput("input_test.txt")
	if err != nil {
		log.Fatal(err)
	}

	ans := Part1(data)
	correct := 2

	if ans != correct {
		t.Errorf("Part 1 = %d, want %d", ans, correct)
	}

}

func TestPart2(t *testing.T) {

	data, err := util.GetInput("input_test.txt")
	if err != nil {
		log.Fatal(err)
	}

	ans := Part2(data)
	correct := 4

	if ans != correct {
		t.Errorf("Part 2 = %d, want %d", ans, correct)
	}

}

func BenchmarkPart1(b *testing.B) {

	data, err := util.GetInput("input_test.txt")
	if err != nil {
		log.Fatal(err)
	}

	for n := 0; n < b.N; n++ {
		Part1(data)
	}

}

func BenchmarkPart2(b *testing.B) {

	data, err := util.GetInput("input_test.txt")
	if err != nil {
		log.Fatal(err)
	}

	for n := 0; n < b.N; n++ {
		Part2(data)
	}

}
