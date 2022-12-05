package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"aoc2022/utility"
)

func main() {
	p, err := filepath.Abs(os.Args[1])
	if err != nil {
		log.Fatal(err.Error())
	}

	text, err := utility.ReadFile(p)
	if err != nil {
		log.Fatal(err.Error())
	}

	textArr := strings.Split(text, "\r\n")

	var pairs [][]Pair
	for _, each := range textArr {
		elves := strings.Split(each, ",")

		elf1 := strings.Split(elves[0], "-")
		px, _ := strconv.Atoi(elf1[0])
		py, _ := strconv.Atoi(elf1[1])
		p1 := NewPair(px, py)

		elf2 := strings.Split(elves[1], "-")
		px, _ = strconv.Atoi(elf2[0])
		py, _ = strconv.Atoi(elf2[1])

		p2 := NewPair(px, py)

		pairs = append(pairs, []Pair{p1, p2})
	}

	var fullyContained int
	for _, p := range pairs {
		if p[0].contains(p[1]) || p[1].contains(p[0]) {
			fullyContained += 1
		}
	}

	fmt.Printf("Part 1 answer: %d\n", fullyContained)

	var overlaps int
	for _, p := range pairs {
		if p[0].overlaps(p[1]) {
			overlaps += 1
		}
	}

	fmt.Printf("Part 2 answer: %d\n", overlaps)
}

type Pair struct {
	start, end int
}

func NewPair(x, y int) Pair {
	return Pair{start: x, end: y}
}

func (p Pair) contains(other Pair) bool {
	return (p.start <= other.start && other.start <= p.end) && (p.start <= other.end && other.end <= p.end)
}

func (p Pair) overlaps(other Pair) bool {
	return (other.start <= p.end && p.end <= other.end) || (p.start <= other.end && other.end <= p.end)
}
