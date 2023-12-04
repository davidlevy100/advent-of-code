package main

import (
	"log"
	"testing"

	util "github.com/davidlevy100/advent-of-code/util"
)

var Part1 = part1
var Part2 = part2

var HasPairs = hasPairs
var HasStraddles = hasStraddles

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

	data, err := util.GetInput("input_test2.txt")
	if err != nil {
		log.Fatal(err)
	}

	ans := Part2(data)
	correct := 2

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

func TestHasPairs(t *testing.T) {

	var tests = []struct {
		s    string
		want bool
	}{
		{"qjhvhtzxzqqjkmpb", true},
		{"xxyxx", true},
		{"uurcxstgmygtbstg", true},
		{"ieodomkazucvgmuy", false},
		{"aaa", false},
	}

	for _, tt := range tests {
		testname := tt.s
		t.Run(testname, func(t *testing.T) {
			ans := HasPairs(tt.s)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}

}

func TestHasStraddles(t *testing.T) {

	var tests = []struct {
		s    string
		want bool
	}{
		{"qjhvhtzxzqqjkmpb", true},
		{"xxyxx", true},
		{"uurcxstgmygtbstg", false},
		{"ieodomkazucvgmuy", true},
	}

	for _, tt := range tests {
		testname := tt.s
		t.Run(testname, func(t *testing.T) {
			ans := HasStraddles(tt.s)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}

}
