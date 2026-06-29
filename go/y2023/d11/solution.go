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
	galaxies   []h.Pt
}

func buildSolution(r io.Reader) *solution {
	lines, _ := h.LinesFromReader(r)
	return &solution{
		input:    lines,
		ans1:     0,
		ans2:     0,
		galaxies: []h.Pt{},
	}
}

func (s *solution) findGalaxiesAndEmpty() ([]bool, []bool) {
	rows := make([]bool, len(s.input))
	cols := make([]bool, len(s.input[0]))
	for i := range rows {
		rows[i] = true
	}
	for i := range cols {
		cols[i] = true
	}

	for i, line := range s.input {
		for j, c := range line {
			if c == '#' {
				s.galaxies = append(s.galaxies, h.Pt{R: i, C: j})
				rows[i] = false
				cols[j] = false
			}
		}
	}

	return rows, cols
}

func (s *solution) expand(emptyRows, emptyCols []bool, factor int) []h.Pt {
	out := make([]h.Pt, len(s.galaxies))
	for i, g := range s.galaxies {
		r := g.R
		c := g.C
		for x := 0; x < g.R; x++ {
			if emptyRows[x] {
				r += factor - 1
			}
		}
		for y := 0; y < g.C; y++ {
			if emptyCols[y] {
				c += factor - 1
			}
		}
		out[i] = h.Pt{R: r, C: c}
	}
	return out
}

func distance(a, b h.Pt) int {
	return h.Abs(a.R-b.R) + h.Abs(a.C-b.C)
}

func total(g []h.Pt) int {
	ans := 0
	for i := range len(g) {
		for j := i + 1; j < len(g); j++ {
			ans += distance(g[i], g[j])
		}
	}
	return ans
}

func (s *solution) run1() {
	rows, cols := s.findGalaxiesAndEmpty()
	newG := s.expand(rows, cols, 2)
	s.ans1 = total(newG)
}

func (s *solution) run2() {
	rows, cols := s.findGalaxiesAndEmpty()
	newG := s.expand(rows, cols, 1000000)
	s.ans2 = total(newG)
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
