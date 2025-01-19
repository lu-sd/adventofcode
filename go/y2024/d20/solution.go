package main

import (
	"adventofcode/utils"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func (s *solution) run1() {
	s.findPath(s.start, 0)
	s.walk(s.start)
}

func (s *solution) findPath(p utils.Pt, step int) {
	if s.GetRune(p) == '#' || s.Seen[p] {
		return
	}
	s.path[p] = step
	if s.GetRune(p) == 'E' {
		return
	}
	s.Seen[p] = true
	for _, dir := range utils.Dir4 {
		nextP := p.PMove(dir)
		s.findPath(nextP, step+1)
	}
	s.Seen[p] = false
}

func (s *solution) walk(p utils.Pt) {
	if s.GetRune(p) == '#' || s.Seen[p] || s.GetRune(p) == 'E' {
		return
	}
	start, cheatEl, find := s.cheat(p)
	if find {
		for _, pt := range cheatEl {
			cheatS := 2
			pathS := s.path[pt] - s.path[start]
			s.saves[pathS-cheatS]++
		}
	}
	s.Seen[p] = true
	for _, dir := range utils.Dir4 {
		nextp := p.PMove(dir)
		s.walk(nextp)
	}
	s.Seen[p] = false
}

func (s *solution) cheat(p utils.Pt) (start utils.Pt, cheatEl []utils.Pt, find bool) {
	for _, dir := range utils.Dir4 {
		cheat := p.Move(dir.C*2, dir.R*2)
		if s.IsInside(cheat) && s.GetRune(cheat) != '#' && !s.Seen[cheat] {
			cheatEl = append(cheatEl, cheat)
		}
	}
	if len(cheatEl) > 0 {
		return p, cheatEl, true
	}
	return p, nil, false
}

func (s *solution) run2() {
}

func (s *solution) res1() int {
	for save, num := range s.saves {
		if save >= 100 {
			s.ans += num
		}
	}
	return s.ans
}

func (s *solution) res2() int {
	return s.ans
}

func buildSolution(r io.Reader) *solution {
	lines, err := utils.LinesFromReader(r)
	if err != nil {
		log.Fatalf("could not read input: %v %v", lines, err)
	}
	var start utils.Pt
	for i, line := range lines {
		for j, rune := range line {
			if rune == 'S' {
				start = utils.Pt{R: i, C: j}
			}
		}
	}

	return &solution{
		ans:   0,
		start: start,
		path:  map[utils.Pt]int{},
		StringGrid: utils.StringGrid{
			NCol:  len(lines[0]),
			NRow:  len(lines),
			Array: lines,
		},
		Seen:  utils.Seen{},
		saves: map[int]int{},
	}
}

type solution struct {
	ans   int
	start utils.Pt
	path  map[utils.Pt]int
	utils.Seen
	utils.StringGrid
	saves map[int]int
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
