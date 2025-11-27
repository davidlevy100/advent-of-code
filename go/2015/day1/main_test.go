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
		{"(())", 0},
		{"(()(()(", 3},
		{"))(((((", 3},
		{"))(", -1},
		{")())())", -3},
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
		{")", 1},
		{"()())", 5},
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
		Part1("))(((((")
	}

}

func BenchmarkPart2(b *testing.B) {

	for n := 0; n < b.N; n++ {
		Part2("))(((((")
	}

}
