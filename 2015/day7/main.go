package main

import (
	"fmt"
	"strconv"
	"strings"

	util "github.com/davidlevy100/advent-of-code/util"
)

type queue struct {
	elements [][]string
}

func (q *queue) enqueue(ss []string) {
	q.elements = append(q.elements, ss)
}

func (q *queue) dequeue() []string {

	result := q.elements[0]
	q.elements = q.elements[1:]

	return result
}

func (q *queue) Len() int {
	return len(q.elements)
}

func main() {

	data, _ := util.GetInput("input.txt")

	fmt.Printf("Part 1 answer: %d\n", part1(data))
	fmt.Printf("Part 2 answer: %d\n", part2(data))

}

func part1(data []string) int {

	var result int

	wires := make(map[string]uint16)
	q := new(queue)

	for _, line := range data {
		msg := strings.Fields(line)
		q.enqueue(msg)
	}

	for q.Len() > 0 {
		parseInput(q.dequeue(), wires, q)
	}

	result = int(wires["a"])

	return result

}

func part2(data []string) int {

	var result int

	wires := make(map[string]uint16)
	wires["b"] = 956
	q := new(queue)

	for _, line := range data {
		msg := strings.Fields(line)
		q.enqueue(msg)
	}

	for q.Len() > 0 {
		parseInput(q.dequeue(), wires, q)
	}

	result = int(wires["a"])

	return result

}

func parseInput(msg []string, m map[string]uint16, q *queue) {

	switch len(msg) {
	case 3:
		assign(msg, m, q)
	case 4:
		complement(msg, m, q)
	case 5:
		bitwise(msg, m, q)
	}

}

func assign(ss []string, m map[string]uint16, q *queue) {

	key := ss[2]

	_, ok := m[key]
	if ok {
		return
	}

	val, err := strconv.ParseUint(ss[0], 10, 16)
	if err != nil {
		_, ok := m[ss[0]]
		if !ok {
			q.enqueue(ss)
		} else {
			m[key] = m[ss[0]]
		}
	} else {
		m[key] = uint16(val)
	}
}

func bitwise(ss []string, m map[string]uint16, q *queue) {

	key := ss[4]
	var op1 uint16
	var op2 uint16

	op, err := strconv.ParseUint(ss[0], 10, 16)
	if err != nil {
		_, ok := m[ss[0]]
		if !ok {
			q.enqueue(ss)
			return
		} else {
			op1 = m[ss[0]]
		}

	} else {
		op1 = uint16(op)
	}

	op, err = strconv.ParseUint(ss[2], 10, 16)
	if err != nil {
		_, ok := m[ss[2]]
		if !ok {
			q.enqueue(ss)
			return
		} else {
			op2 = m[ss[2]]
		}

	} else {
		op2 = uint16(op)
	}

	switch ss[1] {
	case "AND":
		m[key] = op1 & op2
	case "OR":
		m[key] = op1 | op2
	case "LSHIFT":
		m[key] = op1 << op2
	case "RSHIFT":
		m[key] = op1 >> op2
	}

}

func complement(ss []string, m map[string]uint16, q *queue) {

	key := ss[3]
	var op1 uint16
	op, err := strconv.ParseUint(ss[1], 10, 16)
	if err != nil {
		_, ok := m[ss[1]]
		if !ok {
			q.enqueue(ss)
			return
		} else {
			op1 = m[ss[1]]
		}
	} else {
		op1 = uint16(op)
	}

	m[key] = ^op1

}
