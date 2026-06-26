package main

import (
	"adventofcode/h"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type solution struct {
	input      []string
	ans1, ans2 int
}

func buildSolution(r io.Reader) *solution {
	lines, _ := h.LinesFromReader(r)
	return &solution{
		input: lines,
		ans1:  0,
		ans2:  0,
	}
}

type lens struct {
	label string
	focal int
}

func (s *solution) run1() {
	// for _, line := range s.input {
	line := s.input[0]
	for str := range strings.SplitSeq(line, ",") {
		s.ans1 += hash(str)
	}
	// }
}

func hash(s string) int {
	h := 0
	for _, c := range s {
		h += int(c)
		h *= 17
		h = h % 256
	}
	return h
}

func (s *solution) run2() {
	boxes := make([][]lens, 256)

	input := s.input[0]
	for str := range strings.SplitSeq(input, ",") {
		label, focal, hasEqual := findLabel(str)
		putlen(label, focal, hasEqual, boxes)
	}
	for i, box := range boxes {
		for j, v := range box {
			s.ans2 += (i + 1) * (j + 1) * v.focal
		}
	}
}

func putlen(label string, focal int, hasEqual bool, boxes [][]lens) {
	boxNum := hash(label)

	if !hasEqual {
		for i, v := range boxes[boxNum] {
			if v.label == label {
				boxes[boxNum] = append(boxes[boxNum][:i], boxes[boxNum][i+1:]...)
				break
			}
		}
	} else {
		found := false
		for i, v := range boxes[boxNum] {
			if v.label == label {
				boxes[boxNum][i].focal = focal
				found = true
				break
			}
		}
		if !found {
			boxes[boxNum] = append(boxes[boxNum], lens{label, focal})
		}
	}
}

func findLabel(str string) (string, int, bool) {
	if strings.HasSuffix(str, "-") {
		label := str[:len(str)-1]
		return label, 0, false
	} else {
		label, focal, _ := strings.Cut(str, "=")
		focalInt, _ := strconv.Atoi(focal)
		return label, focalInt, true
	}
}

func (s *solution) res1() int {
	return s.ans1
}

func (s *solution) res2() int {
	return s.ans2
}

func part1(r io.Reader) int {
	s := buildSolution(r)
	s.run1()
	return s.res1()
}

func part2(r io.Reader) int {
	s := buildSolution(r)
	s.run2()
	return s.res2()
}

func main() {
	Input, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("fail open input.txt %v", err)
	}
	start := time.Now()
	result := part1(Input)
	elapsed := time.Since(start)
	fmt.Printf("p1 res 🙆-> %d (Time taken: %s)\n", result, elapsed)
	start = time.Now()
	result = part2(Input)
	elapsed = time.Since(start)
	fmt.Printf("p2 res 🙆-> %d (Time taken: %s)\n", result, elapsed)
}
