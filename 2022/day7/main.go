package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	util "github.com/davidlevy100/advent-of-code/util"
)

type file struct {
	name     string
	isDir    bool
	size     int
	parent   *file
	children map[string]*file
}

func newFile(name string, isDir bool) *file {
	return &file{name: name, isDir: isDir, children: make(map[string]*file)}
}

func (f *file) addChild(name string, isDir bool, size int) {
	child := newFile(name, isDir)
	child.size = size
	child.parent = f
	f.children[name] = child
}

func (f *file) cd(s string) (*file, error) {

	if s == "/" {
		newDir := f

		for newDir.parent != nil {
			newDir = newDir.parent
		}

		return newDir, nil

	}

	if s == ".." {

		if f.parent != nil {
			return f.parent, nil
		} else {
			return f, nil
		}

	}

	if result, ok := f.children[s]; ok {

		if result.isDir {
			return result, nil
		} else {
			return nil, fmt.Errorf("not a directory")
		}

	} else {
		return nil, fmt.Errorf("no such file or directory")
	}

}

func main() {

	data, _ := util.GetInput("input.txt")

	fmt.Printf("Part 1 answer: %d\n", part1(data))
	fmt.Printf("Part 2 answer: %d\n", part2(data))

}

func part1(data []string) int {

	var result int

	root := buildTree(data)
	setSizes(root)

	var dfs func(root *file)
	dfs = func(root *file) {
		for _, value := range root.children {
			if value.isDir {
				if value.size < 100000 {
					result += value.size
				}
				dfs(value)
			}
		}
	}

	dfs(root)

	return result

}

func part2(data []string) int {

	var result int

	const total int = 70000000
	const needed int = 30000000

	root := buildTree(data)
	setSizes(root)

	var unused = total - root.size

	var spaceNeeded = needed - unused

	var candidates []int

	var dfs func(root *file)
	dfs = func(root *file) {
		for _, value := range root.children {
			if value.isDir {
				if value.size > spaceNeeded {
					candidates = append(candidates, value.size)
				}
				dfs(value)
			}
		}
	}

	dfs(root)

	sort.Ints(candidates)

	result = candidates[0]

	return result

}

func buildTree(data []string) *file {

	result := newFile("/", true)

	cwd := result

	for _, thisLine := range data {
		thisCommand := strings.Fields(thisLine)

		switch thisCommand[0] {
		case "$":
			if thisCommand[1] == "cd" {
				newDir, err := cwd.cd(thisCommand[2])
				if err != nil {
					fmt.Println(err)
				} else {
					cwd = newDir
				}
			} else {
				continue
			}
		case "dir":
			cwd.addChild(thisCommand[1], true, 0)
		default:
			size, _ := strconv.Atoi(thisCommand[0])
			cwd.addChild(thisCommand[1], false, size)
		}

	}

	return result

}

func setSizes(root *file) {

	for _, value := range root.children {
		if value.isDir {
			setSizes(value)
			root.size += value.size
		} else {
			root.size += value.size
		}
	}
}
