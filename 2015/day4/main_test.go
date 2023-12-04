package main

import (
	"testing"
)

var Part1 = part1
var Part2 = part2

func TestPart1(t *testing.T) {

	var tests = []struct {
		s    string
		want int
	}{
		{"abcdef", 609043},
		{"pqrstuv", 1048970},
	}

	for _, tt := range tests {
		testname := tt.s
		t.Run(testname, func(t *testing.T) {
			ans := Part1(tt.s)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}

}

func TestPart2(t *testing.T) {

	var tests = []struct {
		s    string
		want int
	}{
		{"abcdef", 609043},
		{"pqrstuv", 1048970},
	}

	for _, tt := range tests {
		testname := tt.s
		t.Run(testname, func(t *testing.T) {
			ans := Part2(tt.s)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}

}

func BenchmarkPart1(b *testing.B) {

	for n := 0; n < b.N; n++ {
		Part1("abcdef")
	}

}

func BenchmarkPart2(b *testing.B) {

	for n := 0; n < b.N; n++ {
		Part2("abcdef")
	}

}
