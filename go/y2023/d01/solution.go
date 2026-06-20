package main

import (
	"adventofcode/h"
	"fmt"
	"io"
	"log"
	"os"
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

var m = [...]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func findLnR2(l string) (fst, lst int) {
	fst, lst = -1, -1

	for i := 0; i < len(l); i++ {

		if v, ok := h.IsDigit(rune(l[i])); ok {
			if fst == -1 {
				fst = v
			}
			lst = v
		}

		for j, v := range m {
			if strings.HasSuffix(l[:i+1], v) {
				if fst == -1 {
					fst = j + 1
				}
				lst = j + 1
			}
		}
	}
	return
}

func findLnR(l string) (fst, lst int) {
	fst, lst = -1, -1
	for _, r := range l {
		if v, ok := h.IsDigit(r); ok {
			if fst == -1 {
				fst = v
			}
			lst = v
		}
	}
	return
}

func (s *solution) run1() {
	for _, line := range s.input {
		left, right := findLnR(line)
		s.ans1 += left*10 + right

	}
}

func (s *solution) run2() {
	for _, line := range s.input {
		left, right := findLnR2(line)
		s.ans2 += left*10 + right
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
