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
	cheatEl, find := s.findCheat(p)
	if find {
		for _, end := range cheatEl {
			cheatS := 2
			pathS := s.path[end] - s.path[p]
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

func (s *solution) findCheat(p utils.Pt) (cheatEl []utils.Pt, find bool) {
	for _, dir := range utils.Dir4 {
		cheat := p.Move(dir.C*2, dir.R*2)
		if s.IsInside(cheat) && s.GetRune(cheat) != '#' && !s.Seen[cheat] {
			cheatEl = append(cheatEl, cheat)
		}
	}
	if len(cheatEl) > 0 {
		return cheatEl, true
	}
	return nil, false
}

func (s *solution) run2() {
	s.findPath(s.start, 0)
	s.walk2(s.start)
}

func (s *solution) walk2(p utils.Pt) {
	if s.GetRune(p) == '#' || s.GetRune(p) == 'E' || s.Seen[p] {
		return
	}
	s.findCheat2(p)
	s.Seen[p] = true
	for _, dir := range utils.Dir4 {
		nextP := p.PMove(dir)
		s.walk2(nextP)
	}
	s.Seen[p] = false
}

func (s *solution) findCheat2(start utils.Pt) {
	queue := []utils.Pt{start}
	visited := map[utils.Pt]bool{}
	visited[start] = true
	step := 0
	for len(queue) > 0 {
		n := len(queue)
		for n > 0 {
			n--
			p := queue[0]
			queue = queue[1:]
			if s.GetRune(p) != '#' && p != start {
				dist := s.path[p] - s.path[start]
				if save := dist - step; save >= 100 {
					s.saves[save]++
				}
			}
			for _, dir := range utils.Dir4 {
				nextP := p.PMove(dir)
				if s.IsInside(nextP) && !visited[nextP] {
					visited[nextP] = true
					queue = append(queue, nextP)
				}
			}
		}
		step++
		if step > 20 {
			break
		}
	}
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
	for _, num := range s.saves {
		s.ans += num
	}
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
