package main

import (
	"adventofcode/utils"
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
	lines, _ := utils.LinesFromReader(r)
	return &solution{
		input: lines,
		ans1:  0,
		ans2:  0,
	}
}

func (s *solution) run2() {
	for _, line := range s.input {
		outs := parseInput(line)
		var b, r, g int
		for _, v := range outs {
			b = max(v["blue"], b)
			r = max(v["red"], r)
			g = max(v["green"], g)
		}
		mul := b * r * g
		s.ans2 += mul
	}
}

var maxVal = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func isValid(m map[string]int) bool {
	for c, v := range m {
		if v > maxVal[c] {
			return false
		}
	}
	return true
}

func parseInput(s string) []map[string]int {
	var outs []map[string]int

	for value := range strings.SplitSeq(strings.Split(s, ":")[1], ";") {
		dict := make(map[string]int)

		for cube := range strings.SplitSeq(value, ",") {
			// val := strings.Fields(cube)
			// v, c := val[0], val[1]
			// intV, _ := strconv.Atoi(v)
			var (
				c string
				v int
			)
			fmt.Sscanf(cube, "%d %s", &v, &c)
			dict[c] = v
		}
		outs = append(outs, dict)
	}
	return outs
}

func (s *solution) run1() {
	for id, line := range s.input {
		outs := parseInput(line)

		validGame := true

		for _, out := range outs {
			if !isValid(out) {
				validGame = false
				break
			}
		}

		if validGame {
			s.ans1 += id + 1
		}
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
