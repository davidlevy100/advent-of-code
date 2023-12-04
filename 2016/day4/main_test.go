package main

import (
	"log"
	"testing"

	util "github.com/davidlevy100/advent-of-code/util"
)

var Part1 = part1
var Part2 = part2
var CheckCode = checkCode
var ShiftCypher = shiftCypher

func TestCheckCode(t *testing.T) {

	var tests = []struct {
		s     string
		want  int
		want2 bool
	}{
		{"aaaaa-bbb-z-y-x-123[abxyz]", 123, true},
		{"a-b-c-d-e-f-g-h-987[abcde]", 987, true},
		{"not-a-real-room-404[oarel]", 404, true},
		{"totally-real-room-200[decoy]", 0, false},
	}

	for _, tt := range tests {
		testname := tt.s
		t.Run(testname, func(t *testing.T) {
			ans, ans2 := CheckCode(tt.s)
			if ans != tt.want || ans2 != tt.want2 {
				t.Errorf("got %v, want %v, %v", ans, tt.want, tt.want2)
			}
		})
	}

}

func TestShiftCypher(t *testing.T) {

	var tests = []struct {
		s     string
		want  string
		want2 int
	}{
		{"qzmt-zixmtkozy-ivhz-343[abcd]", "very encrypted name", 343},
	}

	for _, tt := range tests {
		testname := tt.s
		t.Run(testname, func(t *testing.T) {
			ans, ans2 := ShiftCypher(tt.s)
			if ans != tt.want || ans2 != tt.want2 {
				t.Errorf("got %s, %d want %s, %d", ans, ans2, tt.want, tt.want2)
			}
		})
	}

}

func TestPart1(t *testing.T) {

	data, err := util.GetInput("input_test.txt")
	if err != nil {
		log.Fatal(err)
	}

	ans := Part1(data)
	correct := 1514

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
	correct := 1514

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
