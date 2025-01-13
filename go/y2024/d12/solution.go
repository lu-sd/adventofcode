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
	for i, line := range s.Array {
		for j, rune := range line {
			cp := utils.Pt{C: j, R: i}
			if s.Seen[cp] {
				continue
			}
			s.dfs(cp, rune, cp)
		}
	}
}

func (s *solution) dfs(curP utils.Pt, flower rune, head utils.Pt) {
	if !s.IsInside(curP) || s.GetRune(curP) != flower || s.Seen[curP] {
		return
	}
	s.Seen[curP] = true
	s.land[head] = append(s.land[head], curP)
	for _, dir := range utils.Dir4 {
		nextP := curP.PMove(dir)
		s.dfs(nextP, flower, head)
	}
}

func (s *solution) run2() {
}

func (s *solution) res1() int {
	for head, lands := range s.land {
		area := len(lands)
		flower := s.GetRune(head)
		prem := s.getPrem(lands, flower)
		s.ans += area * prem
	}
	return s.ans
}

func (s *solution) getPrem(lands []utils.Pt, flower rune) int {
	p := 0
	for _, land := range lands {
		for _, dir := range utils.Dir4 {
			nextP := land.PMove(dir)
			if !s.IsInside(nextP) || s.GetRune(nextP) != flower {
				p++
			}
		}
	}
	return p
}

func (s *solution) res2() int {
	return s.ans
}

func buildSolution(r io.Reader) *solution {
	lines, err := utils.LinesFromReader(r)
	if err != nil {
		log.Fatalf("could not read input: %v %v", lines, err)
	}
	nrow, ncol := len(lines), len(lines[0])
	array := []string{}
	for _, line := range lines {
		array = append(array, line)
	}
	return &solution{
		ans: 0,
		StringGrid: utils.StringGrid{
			NCol:  ncol,
			NRow:  nrow,
			Array: array,
		},
		Seen: utils.Seen{},
		land: map[utils.Pt][]utils.Pt{},
	}
}

type solution struct {
	ans  int
	land map[utils.Pt][]utils.Pt
	utils.StringGrid
	utils.Seen
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
