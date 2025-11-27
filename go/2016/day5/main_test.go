package main

import (
	"testing"
)

var Part1 = part1
var Part2 = part2

func TestPart1(t *testing.T) {

	var tests = []struct {
		input string
		want  string
	}{
		{"abc", "18f47a30"},
	}

	for _, tt := range tests {
		testname := tt.input
		t.Run(testname, func(t *testing.T) {
			ans := Part1(tt.input)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}

}

func TestPart2(t *testing.T) {

	var tests = []struct {
		input string
		want  string
	}{
		{"abc", "05ace8e3"},
	}

	for _, tt := range tests {
		testname := tt.input
		t.Run(testname, func(t *testing.T) {
			ans := Part2(tt.input)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}

func BenchmarkPart1(b *testing.B) {

	for n := 0; n < b.N; n++ {
		Part1("abc")
	}

}

func BenchmarkPart2(b *testing.B) {

	for n := 0; n < b.N; n++ {
		Part2("abc")
	}

}
