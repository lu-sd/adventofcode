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
	grid                   h.StringGrid
	ans1, ans2, nrow, ncol int
	gearM                  map[h.Pt][]int
}

func buildSolution(r io.Reader) *solution {
	lines, _ := h.LinesFromReader(r)
	return &solution{
		grid:  h.StringGrid{Array: lines},
		ans1:  0,
		ans2:  0,
		nrow:  len(lines),
		ncol:  len(lines[0]),
		gearM: map[h.Pt][]int{},
	}
}

func (s *solution) findValid(p h.Pt) int {
	for _, d := range h.Dir8 {
		nexp := p.PMove(d)
		if !s.grid.IsInside(nexp) {
			continue
		}
		b := s.grid.GetByte(nexp)
		if !h.IsDigitBool(b) && b != '.' {
			return 1
		}
	}
	return 0
}

func (s *solution) findStar(p h.Pt) (h.Pt, bool) {
	for _, d := range h.Dir8 {
		nexp := p.PMove(d)
		if !s.grid.IsInside(nexp) {
			continue
		}
		b := s.grid.GetByte(nexp)
		if b == '*' {
			return nexp, true
		}
	}
	return h.Pt{}, false
}

func (s *solution) tour() {
	for i := 0; i < s.nrow; i++ {
		for j := 0; j < s.ncol; j++ {
			curP := h.Pt{R: i, C: j}
			b := s.grid.GetByte(curP)

			if h.IsDigitBool(b) {
				var num, valid int
				localM := map[h.Pt]bool{}

				for h.IsDigitBool(b) {
					valid += s.findValid(curP)
					// find all star around this number
					if gear, ok := s.findStar(curP); ok {
						localM[gear] = true
					}
					v, _ := h.IsDigit(b)
					num = num*10 + v
					j++
					curP = h.Pt{R: i, C: j}
					if !s.grid.IsInside(curP) {
						break
					}
					b = s.grid.GetByte(curP)
				}
				for k := range localM {
					s.gearM[k] = append(s.gearM[k], num) // regiest number to all its associated star
				}
				if valid > 0 {
					s.ans1 += num
				}
			}
		}
	}
}

func (s *solution) run1() {
	s.tour()
}

func (s *solution) run2() {
	s.tour()
	for _, v := range s.gearM {
		if len(v) == 2 {
			s.ans2 += v[0] * v[1]
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
