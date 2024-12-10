package main

import (
	"adventofcode/utils"
	"fmt"
	"io"
	"os"
)

type (
	pt       struct{ x, y int }
	head     pt
	tail     pt
	seen     map[pt]bool
	trailend map[head]map[tail]struct{}
	solution struct {
		nrow, ncol int
		seen       seen
		ans        int
		trailGrid  [][]int
		trailend   trailend
	}
)

func (p1 pt) move(p2 pt) pt {
	return pt{p1.x + p2.x, p1.y + p2.y}
}

var dirs = []pt{
	{-1, 0},
	{1, 0},
	{0, 1},
	{0, -1},
}

func (s *solution) isInside(p pt) bool {
	return p.x >= 0 && p.x < s.nrow && p.y >= 0 && p.y < s.ncol
}

func (s *solution) findNine(curPt, headPt pt, target int) {
	if !s.isInside(curPt) {
		return
	}
	if s.trailGrid[curPt.x][curPt.y] != target {
		return
	}
	if s.seen[curPt] {
		return
	}

	if target == 9 {
		hp, tp := head(headPt), tail(curPt)
		if s.trailend[hp] == nil {
			s.trailend[hp] = map[tail]struct{}{}
		}
		s.trailend[hp][tp] = struct{}{}
		return
	}
	for _, dir := range dirs {
		nxtP := curPt.move(dir)
		s.seen[curPt] = true
		s.findNine(nxtP, headPt, target+1)
		s.seen[curPt] = false
	}
}

func (s *solution) findNine2(curPt pt, target int) {
	if !s.isInside(curPt) {
		return
	}
	if s.trailGrid[curPt.x][curPt.y] != target {
		return
	}
	if s.seen[curPt] {
		return
	}

	if target == 9 {
		s.ans++
		return
	}
	for _, dir := range dirs {
		nxtP := curPt.move(dir)
		s.seen[curPt] = true
		s.findNine2(nxtP, target+1)
		s.seen[curPt] = false
	}
}

func (s *solution) run1() {
	for i := range s.nrow {
		for j := range s.ncol {
			if s.trailGrid[i][j] == 0 {
				s.findNine(pt{i, j}, pt{i, j}, 0)
			}
		}
	}
}

func (s *solution) run2() {
	for i := range s.nrow {
		for j := range s.ncol {
			if s.trailGrid[i][j] == 0 {
				s.findNine2(pt{i, j}, 0)
			}
		}
	}
}

func (s *solution) res() int {
	for _, v := range s.trailend {
		s.ans += len(v)
	}
	return s.ans
}

func buildSolution(r io.Reader) *solution {
	lines, err := utils.LinesFromReader(r)
	if err != nil {
		return nil
	}
	nrow, ncol := len(lines), len(lines[0])
	grid := make([][]int, nrow)
	for i, line := range lines {
		grid[i] = make([]int, ncol)
		for j, c := range line {
			// num, err := strconv.Atoi(string(c))
			// if err != nil {
			// 	fmt.Println("int convert error")
			// 	return nil
			// }
			grid[i][j] = int(c - '0')
		}
	}
	fmt.Println(nrow, ncol)

	return &solution{
		nrow:      nrow,
		ncol:      ncol,
		ans:       0,
		trailGrid: grid,
		trailend:  trailend{},
		seen:      seen{},
	}
}

func part1(r io.Reader) int {
	s := buildSolution(r)
	s.run1()
	return s.res()
}

func part2(r io.Reader) int {
	s := buildSolution(r)
	s.run2()
	return s.ans
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
