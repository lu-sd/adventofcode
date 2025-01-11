package main

import (
	"adventofcode/utils"
	"fmt"
	"io"
	"log"
	"os"
)

type solution struct {
	ans int
	utils.StringGrid
	utils.Seen
	trail map[string]bool
}

func (s *solution) run1() {
	for i, line := range s.Array {
		for j, rune := range line {
			if rune == '0' {
				pt := utils.Pt{C: j, R: i}
				s.dfs1(pt, pt, 0)
			}
		}
	}
}

func (s *solution) dfs1(start, curPt utils.Pt, target int) {
	if !s.IsInside(curPt) || s.Seen[curPt] || s.StringGrid.GetInt(curPt) != target {
		return
	}
	if s.StringGrid.GetInt(curPt) == 9 {
		s.find(start, curPt)
		return
	}

	s.Seen[curPt] = true
	for _, dir := range utils.Dir4 {
		nPt := curPt.PMove(dir)
		s.dfs1(start, nPt, target+1)
	}
	s.Seen[curPt] = false
}

func (s *solution) find(start, end utils.Pt) {
	id := fmt.Sprintf("%d,%d,%d,%d", start.C, start.R, end.C, end.R)
	s.trail[id] = true
}

func (s *solution) run2() {
	for i, line := range s.Array {
		for j, rune := range line {
			if rune == '0' {
				pt := utils.Pt{C: j, R: i}
				s.dfs2(pt, 0)
			}
		}
	}
}

func (s *solution) dfs2(curPt utils.Pt, target int) {
	if !s.IsInside(curPt) || s.Seen[curPt] || s.StringGrid.GetInt(curPt) != target {
		return
	}
	if s.StringGrid.GetInt(curPt) == 9 {
		s.ans++
		return
	}

	s.Seen[curPt] = true
	for _, dir := range utils.Dir4 {
		nPt := curPt.PMove(dir)
		s.dfs2(nPt, target+1)
	}
	s.Seen[curPt] = false
}

func (s *solution) res2() int {
	return s.ans
}

func (s *solution) res1() int {
	return len(s.trail)
}

func buildSolution(r io.Reader) *solution {
	lines, err := utils.LinesFromReader(r)
	if err != nil {
		log.Fatalf("could not read input: %v %v", lines, err)
	}
	nrow := len(lines)
	ncol := len(lines[0])

	return &solution{
		ans: 0,
		StringGrid: utils.StringGrid{
			NRow:  nrow,
			NCol:  ncol,
			Array: lines,
		},
		Seen:  utils.Seen{},
		trail: map[string]bool{},
	}
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
	arg := os.Args[1]
	fmt.Println("Running part", arg)
	switch arg {
	case "1":
		fmt.Println("p1 res 🙆-> ", part1(os.Stdin))
	case "2":
		fmt.Println("p2 res 🙆-> ", part2(os.Stdin))
	}
}
