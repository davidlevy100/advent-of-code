package main

import (
	"testing"
)

var Part1 = part1
var Part2 = part2

func TestPart1(t *testing.T) {

	var tests = []struct {
		s    string
		want bool
	}{
		{"aaaaa-bbb-z-y-x-123[abxyz]", true},
		{"a-b-c-d-e-f-g-h-987[abcde]", true},
		{"not-a-real-room-404[oarel]", true},
		{"totally-real-room-200[decoy]", false},
	}

	for _, tt := range tests {
		testname := tt.s
		t.Run(testname, func(t *testing.T) {
			ans := Part1(tt.s)
			if ans != tt.want {
				t.Errorf("got %b, want %b", ans, tt.want)
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
