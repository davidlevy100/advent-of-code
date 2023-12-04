package main

import (
	"log"
	"testing"

	util "github.com/davidlevy100/advent-of-code/util"
)

var Part1 = part1
var Part2 = part2
var IsABBA = isABBA
var HasABBA = hasAbba
var AbbaTest = abbaTest

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

func TestIsAbba(t *testing.T) {

	var tests = []struct {
		s    string
		want bool
	}{
		{"abba", true},
		{"aaaa", false}, // can't all be the same
		{"mnop", false},
		{"qwer", false},
		{"a", false}, // not 4 letters
		{"xyyx", true},
	}

	for _, tt := range tests {
		testname := tt.s
		t.Run(testname, func(t *testing.T) {
			ans := IsABBA(tt.s)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}

}

func TestHasAbba(t *testing.T) {

	var tests = []struct {
		s    string
		want bool
	}{
		{"abba", true},
		{"aaaaaa", false}, // can't all be the same
		{"habbah", true},
		{"lmnop", false},
		{"qwer", false},
		{"a", false}, //not 4 letters
		{"xyyx", true},
	}

	for _, tt := range tests {
		testname := tt.s
		t.Run(testname, func(t *testing.T) {
			ans := HasABBA(tt.s)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestAbbaTest(t *testing.T) {

	var tests = []struct {
		s    string
		want bool
	}{
		{"abba[mnop]qrst", true},
		{"abcd[bddb]xyyx", false},
		{"aaaa[qwer]tyui", false},
		{"ioxxoj[asdfgh]zxcvbn", true},
	}

	for _, tt := range tests {
		testname := tt.s
		t.Run(testname, func(t *testing.T) {
			ans := AbbaTest(tt.s)
			if ans != tt.want {
				t.Errorf("%s: got %v, want %v", testname, ans, tt.want)
			}
		})
	}
}
