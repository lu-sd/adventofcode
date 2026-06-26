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

func buildSolution(r io.Reader) *solution {
	lines, _ := h.LinesFromReader(r)
	return &solution{
		input: lines,
		ans1:  0,
		ans2:  0,
	}
}

func allZero(data []int) bool {
	for _, v := range data {
		if v != 0 {
			return false
		}
	}
	return true
}

func findDif(d []int) (finnal [][]int) {
	finnal = append(finnal, d)
	for !allZero(d) {
		var dif []int
		for i := 0; i < len(d)-1; i++ {
			dif = append(dif, d[i+1]-d[i])
		}
		finnal = append(finnal, dif)
		d = dif
	}
	return
}

func (s *solution) run1() {
	for _, line := range s.input {
		data := h.IntsFromString(line)
		finnal := findDif(data)
		for _, v := range finnal {
			s.ans1 += v[len(v)-1]
		}

	}
}

func (s *solution) run2() {
	for _, line := range s.input {
		data := h.IntsFromString(line)
		finnal := findDif(data)
		var res []int
		for _, data := range finnal {
			res = append(res, data[0])
		}
		for i, v := range res {
			if i%2 != 0 {
				res[i] = -v
			}
			s.ans2 += res[i]
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
