package main

import (
	"adventofcode/h"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type solution struct {
	input      []string
	ans1, ans2 int
}

type m map[byte]int

func buildSolution(r io.Reader) *solution {
	lines, _ := h.LinesFromReader(r)
	return &solution{
		input: lines,
		ans1:  0,
		ans2:  0,
	}
}

func getFre(l string) rune {
	fre := m{}
	for i := 0; i < len(l)/2; i++ {
		fre[l[i]]++
	}
	for j := len(l) / 2; j < len(l); j++ {
		if fre[l[j]] != 0 {
			return rune(l[j])
		}
	}
	return 0
}

func getFre2(a, b, c string) rune {
	fre := make(map[rune]int)
	com := make(map[rune]int)

	for _, r := range a {
		fre[r]++
	}

	for _, r := range b {
		if fre[r] != 0 {
			com[r]++
		}
	}
	for _, r := range c {
		if com[r] != 0 {
			return r
		}
	}
	return 0
}

func getPrio(r rune) int {
	if r > 'Z' {
		return int(r-'a') + 1
	}
	return int(r-'A') + 27
}

func (s *solution) run2() {
	for i := 0; i < len(s.input); i += 3 {
		fre := getFre2(s.input[i], s.input[i+1], s.input[i+2])
		s.ans2 += getPrio(fre)
	}
}

func (s *solution) run1() {
	for _, line := range s.input {
		fre := getFre(line)
		s.ans1 += getPrio(fre)
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
