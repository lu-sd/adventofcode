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
		ans1:  1,
		ans2:  0,
	}
}

func findWays(t, d int) int {
	var n int
	for i := range t {
		if (t-i)*i > d {
			n++
		}
	}
	return n
}

func TnD(s []string) (t, d string) {
	_, time, _ := strings.Cut(s[0], ":")
	_, dis, _ := strings.Cut(s[1], ":")
	return time, dis
}

func (s *solution) run1() {
	time, dis := TnD(s.input)
	timeInt := h.IntsFromString(time)
	disInt := h.IntsFromString(dis)
	for i := range len(timeInt) {
		s.ans1 *= findWays(timeInt[i], disInt[i])
	}
}

func (s *solution) run2() {
	time, dis := TnD(s.input)
	timeInt, _ := strconv.Atoi(strings.ReplaceAll(time, " ", ""))
	disInt, _ := strconv.Atoi(strings.ReplaceAll(dis, " ", ""))
	s.ans2 = findWays(timeInt, disInt)
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
