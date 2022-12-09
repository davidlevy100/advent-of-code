package main

import (
	"fmt"
	"strconv"
	"strings"

	util "github.com/davidlevy100/advent-of-code/util"
)

type coordinates struct {
	x, y int
}

type knot struct {
	next       *knot
	position   coordinates
	posHistory map[coordinates]int
}

func newKnot() *knot {
	return &knot{position: coordinates{0, 0}, posHistory: map[coordinates]int{{0, 0}: 1}}
}

func (k *knot) move(dir string) {
	switch dir {
	case "U":
		k.position.y++
	case "UL":
		k.position.x--
		k.position.y++
	case "UR":
		k.position.x++
		k.position.y++
	case "D":
		k.position.y--
	case "DL":
		k.position.x--
		k.position.y--
	case "DR":
		k.position.x++
		k.position.y--
	case "L":
		k.position.x--
	case "R":
		k.position.x++
	}

	k.posHistory[k.position]++
}

func (k *knot) follow(other *knot) {
	xDiff := other.position.x - k.position.x
	yDiff := other.position.y - k.position.y

	if abs(xDiff) > 1 || abs(yDiff) > 1 {

		angle := k.getAngle(xDiff, yDiff)
		k.move(angle)

	}
}

func (k *knot) getAngle(xDiff, yDiff int) string {

	var result string

	if yDiff > 0 {
		result += "U"
	} else if yDiff < 0 {
		result += "D"
	}

	if xDiff > 0 {
		result += "R"
	} else if xDiff < 0 {
		result += "L"
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		x *= -1
	}

	return x
}

func main() {

	data, _ := util.GetInput("input.txt")

	fmt.Printf("Part 1 answer: %d\n", part1(data))
	fmt.Printf("Part 2 answer: %d\n", part2(data))

}

func part1(data []string) int {

	return simulation(data, 1)

}

func part2(data []string) int {

	return simulation(data, 9)

}

func simulation(data []string, length int) int {

	var result int

	head := newKnot()
	var this = head
	var tail *knot

	for i := 1; i <= length; i++ {
		this.next = newKnot()
		this = this.next

		if i == length {
			tail = this
		}

	}

	for _, thisLine := range data {
		dir, mag := parseCommand(thisLine)

		for i := 0; i < mag; i++ {
			head.move(dir)
			prev, this := head, head

			for this.next != nil {
				this = this.next
				this.follow(prev)
				prev = prev.next
			}
		}
	}
	result = len(tail.posHistory)

	return result

}

func parseCommand(s string) (string, int) {

	var dir string
	var mag int

	inputs := strings.Fields(s)

	dir = inputs[0]
	mag, _ = strconv.Atoi(inputs[1])

	return dir, mag

}
