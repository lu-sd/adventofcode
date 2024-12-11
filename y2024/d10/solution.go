package main

import (
	"adventofcode/utils"
	"fmt"
	"io"
	"log"
	"os"
)
const a = "fdaf"

type solution struct {
	grid         []string
	nrow, ncol   int
	seen         [][]bool
	ans          int
	dirs         []utils.Pt
	trailHeader  map[utils.Pt]map[utils.Pt]bool
	trailHeader2 map[utils.Pt]map[utils.Pt]int
}


func (s *solution) isInside(pt utils.Pt) bool {
	return pt.X >= 0 && pt.X < s.nrow && pt.Y >= 0 && pt.Y < s.ncol
}

func (s *solution) isSeen(pt utils.Pt) bool {
	return s.seen[pt.X][pt.Y]
}

func (s *solution) pByte(pt utils.Pt) byte {
	return s.grid[pt.X][pt.Y]
}

func (s *solution) pInt(pt utils.Pt) int {
	return int(s.pByte(pt) - '0')
}

func (s *solution) run1() {
	var dfs func(s, c utils.Pt, h int)

	dfs = func(sp, cp utils.Pt, preHeight int) {
		if s.isSeen(cp) || s.pInt(cp)-preHeight != 1 {
			return
		}
		// fmt.Println(sp, cp, "reach ")
		if s.pInt(cp) == 9 {
			if s.trailHeader[sp] == nil {
				s.trailHeader[sp] = map[utils.Pt]bool{}
			}
			s.trailHeader[sp][cp] = true
		}
		for _, dir := range s.dirs {
			nxP := cp.Move(dir.X, dir.Y)
			if !s.isInside(nxP) {
				continue
			}
			s.seen[cp.X][cp.Y] = true
			dfs(sp, nxP, s.pInt(cp))
			s.seen[cp.X][cp.Y] = false
		}
	}

	for i, line := range s.grid {
		for j, r := range line {
			if r == '0' {
				cp := utils.Pt{X: i, Y: j}
				dfs(cp, cp, -1)
			}
		}
	}
}

func (s *solution) run2() {
	var dfs func(s, c utils.Pt, h int)

	dfs = func(sp, cp utils.Pt, preHeight int) {
		if s.isSeen(cp) || s.pInt(cp)-preHeight != 1 {
			return
		}
		// fmt.Println(sp, cp, "reach ")
		if s.pInt(cp) == 9 {
			if s.trailHeader2[sp] == nil {
				s.trailHeader2[sp] = map[utils.Pt]int{}
			}
			s.trailHeader2[sp][cp]++
		}
		for _, dir := range s.dirs {
			nxP := cp.Move(dir.X, dir.Y)
			if !s.isInside(nxP) {
				continue
			}
			s.seen[cp.X][cp.Y] = true
			dfs(sp, nxP, s.pInt(cp))
			s.seen[cp.X][cp.Y] = false
		}
	}

	for i, line := range s.grid {
		for j, r := range line {
			if r == '0' {
				cp := utils.Pt{X: i, Y: j}
				dfs(cp, cp, -1)
			}
		}
	}
}

func (s *solution) res() int {
	for _, v := range s.trailHeader {
		s.ans += len(v)
	}
	return s.ans
}

func (s *solution) res2() int {
	for _, v := range s.trailHeader2 {
		for _, v := range v {
			s.ans += v
		}
	}
	return s.ans
}

func buildSolution(r io.Reader) *solution {
	lines, err := utils.LinesFromReader(r)
	nr, nc := len(lines), len(lines[0])
	if err != nil {
		log.Fatalf("could not read input: %v %v", lines, err)
	}
	seen := make([][]bool, nr)
	for i := range nr {
		seen[i] = make([]bool, nc)
	}
	dirs := []utils.Pt{
		{X: -1, Y: 0},
		{X: 1, Y: 0},
		{X: 0, Y: 1},
		{X: 0, Y: -1},
	}

	return &solution{
		nrow:         nr,
		ncol:         nc,
		seen:         seen,
		dirs:         dirs,
		grid:         lines,
		trailHeader:  make(map[utils.Pt]map[utils.Pt]bool),
		trailHeader2: make(map[utils.Pt]map[utils.Pt]int),
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
