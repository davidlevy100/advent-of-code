package main

import (
	"fmt"
	"strconv"
	"strings"

	util "github.com/davidlevy100/advent-of-code/util"
)

type coords struct {
	x, y int
}

type sleigh struct {
	dir int
	pos coords
	loc map[coords]int
}

func newSleigh() *sleigh {

	s := new(sleigh)
	s.loc = make(map[coords]int)
	s.loc[coords{0, 0}] = 1

	return s
}

func (s *sleigh) setRotationAndGetMagnitude(c string) int {

	dir := c[0]
	m := c[1:]
	mag, _ := strconv.Atoi(m)

	switch dir {
	case 'R':
		s.dir += 90
	case 'L':
		s.dir -= 90
	}

	if s.dir < 0 {
		s.dir = 360 + s.dir
	}

	s.dir = s.dir % 360

	return mag

}

func (s *sleigh) move() coords {

	var result coords

	switch s.dir {
	case 0:
		s.pos.y++
	case 90:
		s.pos.x++
	case 180:
		s.pos.y--
	case 270:
		s.pos.x--
	}

	result = coords{s.pos.x, s.pos.y}
	s.loc[result]++
	return result
}

func (s *sleigh) dist(x2, y2 int) int {

	return abs(s.pos.x-x2) + abs(s.pos.y-y2)

}

func abs(i int) int {
	if i < 0 {
		i *= -1
	}
	return i
}

func main() {

	data, _ := util.GetInput("input.txt")

	fmt.Printf("Part 1 answer: %d\n", part1(data[0]))
	fmt.Printf("Part 2 answer: %d\n", part2(data[0]))

}

func part1(data string) int {

	var result int

	commands := strings.Split(data, ", ")

	s := newSleigh()

	for _, c := range commands {
		mag := s.setRotationAndGetMagnitude(c)
		for i := 0; i < mag; i++ {
			s.move()
		}
	}

	result = s.dist(0, 0)

	return result

}

func part2(data string) int {

	var result int

	commands := strings.Split(data, ", ")

	s := newSleigh()

LOOP:
	for _, c := range commands {
		mag := s.setRotationAndGetMagnitude(c)
		for i := 0; i < mag; i++ {
			theseCoords := s.move()
			if s.loc[theseCoords] == 2 {
				break LOOP
			}
		}
	}

	result = s.dist(0, 0)
	return result
}
