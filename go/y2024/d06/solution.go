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
	seen     map[utils.Pt]bool
	spotGrid [][]int
	start    utils.Pt
}

var turns = [4][2]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func (s *solution) run1() {
	for r, line := range s.Array {
		for c, rune := range line {
			if rune == '^' {
				s.dfs(r, c, 0)
				return
			}
		}
	}
}

func (s *solution) dfs(row, col, dir int) {
	pt := utils.Pt{C: col, R: row}
	s.seen[pt] = true
	curDir := turns[dir%4]
	nr, nc := row+curDir[0], col+curDir[1]
	Npt := utils.Pt{C: nc, R: nr}
	if !s.IsInside(Npt) {
		return
	}
	if s.GetRune(Npt) == '#' {
		s.dfs(row, col, dir+1)
	} else {
		s.dfs(nr, nc, dir)
	}
}

func (s *solution) run2() int {
	res := 0
	for r, line := range s.Array {
		for c := range line {
			if s.spotGrid[r][c] == 1 {
				seen := map[string]bool{}
				s.spotGrid[r][c] = -1
				res += s.isInfinity(s.start, 0, seen)
				s.spotGrid[r][c] = 1

			}
		}
	}
	return res
}

func (s *solution) isInfinity(pt utils.Pt, dir int, seen map[string]bool) int {
	curdir := turns[dir%4]
	nr, nc := pt.R+curdir[0], pt.C+curdir[1]
	Npt := utils.Pt{
		R: nr,
		C: nc,
	}
	if !s.IsInside(Npt) {
		return 0
	}
	if s.GetRune(Npt) != '#' && s.spotGrid[nr][nc] != -1 {
		return s.isInfinity(Npt, dir, seen)
	} else {
		id := fmt.Sprintf("%d,%d,%d", nr, nc, dir%4)
		if seen[id] {
			return 1
		}
		seen[id] = true
		return s.isInfinity(pt, dir+1, seen)
	}
}

func (s *solution) res() int {
	for _, v := range s.seen {
		if v {
			s.ans++
		}
	}
	return s.ans
}

func buildSolution(r io.Reader) *solution {
	lines, err := utils.LinesFromReader(r)
	if err != nil {
		log.Fatalf("could not read input: %v %v", lines, err)
	}
	Array := []string{}
	// spotGrid := [][]int{}
	spotGrid := make([][]int, len(lines))
	var start utils.Pt
	for r, line := range lines {
		spotGrid[r] = make([]int, len(line))
		Array = append(Array, line)
		for c, rune := range line {
			if rune == '.' {
				spotGrid[r][c] = 1
			}
			if rune == '^' {
				start = utils.Pt{C: c, R: r}
			}
		}
	}

	return &solution{
		ans: 0,
		StringGrid: utils.StringGrid{
			Array: Array,
		},
		seen:     map[utils.Pt]bool{},
		start:    start,
		spotGrid: spotGrid,
	}
}

func part1(r io.Reader) int {
	s := buildSolution(r)
	s.run1()
	return s.res()
}

func part2(r io.Reader) int {
	s := buildSolution(r)
	return s.run2()
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
