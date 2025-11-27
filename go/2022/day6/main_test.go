package main

import (
	"fmt"
	"testing"
)

var Part1 = part1
var Part2 = part2

func TestPart1(t *testing.T) {

	var tests = []struct {
		s    string
		want int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
		{"nppdvjthqldpwncqszvftbrmjlhg", 6},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.s)
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
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 23},
		{"nppdvjthqldpwncqszvftbrmjlhg", 23},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.s)
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
		Part1("mjqjpqmgbljsphdztnvjfqwrcgsmlb")
	}

}

func BenchmarkPart2(b *testing.B) {

	for n := 0; n < b.N; n++ {
		Part2("mjqjpqmgbljsphdztnvjfqwrcgsmlb")
	}

}
