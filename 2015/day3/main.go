package main

import (
	"fmt"

	util "github.com/davidlevy100/advent-of-code/util"
)

type done struct{}

type coords struct {
	x, y int
}

type santa struct {
	position coords
}

func (s *santa) move(dir rune, visited map[coords]done) {

	switch dir {
	case '^':
		s.position.y++
	case 'v':
		s.position.y--
	case '<':
		s.position.x--
	case '>':
		s.position.x++
	}

	visited[s.position] = done{}

}

func newSanta() *santa {

	result := &santa{}
	return result

}

func main() {

	data, _ := util.GetInput("input.txt")

	fmt.Printf("Part 1 answer: %d\n", part1(data[0]))
	fmt.Printf("Part 2 answer: %d\n", part2(data[0]))

}

func part1(data string) int {

	var result int
	var visited = make(map[coords]done)
	visited[coords{}] = done{}

	thisSanta := *newSanta()

	for _, thisDir := range data {
		thisSanta.move(thisDir, visited)
	}

	result = len(visited)

	return result

}

func part2(data string) int {

	var result int

	thisSanta := newSanta()
	roboSanta := newSanta()

	var visited = make(map[coords]done)
	visited[coords{}] = done{}

	var roboTurn bool

	for _, thisDir := range data {

		if roboTurn {
			roboSanta.move(thisDir, visited)
		} else {
			thisSanta.move(thisDir, visited)
		}

		roboTurn = !roboTurn

	}

	result = len(visited)

	return result

}
