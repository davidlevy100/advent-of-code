package main

import (
	"sort"
	"strconv"
	"strings"
)

type pair struct {
	key   rune
	value int
}

type pairList []pair

func (p pairList) Len() int      { return len(p) }
func (p pairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p pairList) Less(i, j int) bool {

	if p[i].value == p[j].value {
		return p[i].key < p[j].key
	}

	return p[i].value > p[j].value
}

func checkCode(s string) (int, bool) {

	var result bool = true

	count := make(map[rune]int)

	codeParts := strings.Split(s, "-")

	last := strings.Split(codeParts[len(codeParts)-1], "[")

	sectorID, _ := strconv.Atoi(last[0])
	checksum := last[1][:len(last[1])-1]

	for _, thisPart := range codeParts[:len(codeParts)-1] {
		for _, thisRune := range thisPart {
			count[thisRune]++
		}
	}

	p := make(pairList, len(count))

	i := 0
	for k, v := range count {
		p[i] = pair{k, v}
		i++
	}

	sort.Sort(p)

	for i, thisRune := range checksum {

		if thisRune != p[i].key {
			return 0, false
		}
	}

	return sectorID, result

}

func shiftCypher(s string) (string, int) {

	var result string

	codeParts := strings.Split(s, "-")

	last := strings.Split(codeParts[len(codeParts)-1], "[")

	sectorID, _ := strconv.Atoi(last[0])

	var words []string

	for _, thisPart := range codeParts[:len(codeParts)-1] {

		var word string
		for _, thisRune := range thisPart {
			newLetter := string(rune((int(thisRune-'a')+sectorID)%26) + 'a')
			word += newLetter
		}

		words = append(words, word)

	}

	result = strings.Join(words, " ")

	return result, sectorID

}
