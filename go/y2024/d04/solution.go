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
	target string
	dirs   []utils.Pt
}

// var dirs = [
// [1,0],
// [-1,0],
// [0,1],
// [0,-1],
// ] js not go

func (s *solution) hasTarget(row, col int, dir utils.Pt) int {
	for step, rune := range s.target {
		nr, nc := row+step*dir.R, col+step*dir.C
		nextPt := utils.Pt{R: nr, C: nc}
		if !s.IsInside(nextPt) || s.GetRune(nextPt) != rune {
			return 0
		}
	}
	return 1
}

func (s *solution) run1() {
	for r := range s.NRow {
		for c := range s.NCol {
			for _, dir := range s.dirs {
				s.ans += s.hasTarget(r, c, dir)
			}
		}
	}
}

var dirs2 = [2][2][2]int{
	{
		{1, 1},
		{-1, -1},
	},
	{
		{-1, 1},
		{1, -1},
	},
}

func (s *solution) run2() int {
	valueM := map[rune]int{
		'M': 1,
		'S': 2,
	}
	res := 0
	for r, line := range s.Array {
	out:
		for c, rune := range line {
			if rune != 'A' {
				continue
			}
			for _, dir := range dirs2 {
				counter := 0
				for _, point := range dir {
					nr, nc := r+point[0], c+point[1]
					nextP := utils.Pt{R: nr, C: nc}
					if !s.IsInside(nextP) {
						continue out
					}
					counter += valueM[s.GetRune(nextP)]
				}
				if counter != 3 {
					continue out
				}
			}
			res++
		}
	}
	return res
}

func (s *solution) res() int {
	return s.ans
}

func buildSolution(r io.Reader) *solution {
	lines, err := utils.LinesFromReader(r)
	if err != nil {
		log.Fatalf("couldnot read input: %v %v", lines, err)
	}
	arr := []string{}
	for _, line := range lines {
		arr = append(arr, line)
	}
	var dirs []utils.Pt
	slice := []int{0, 1, -1}
	for _, v1 := range slice {
		for _, v2 := range slice {
			if v1 == 0 && v2 == 0 {
				continue
			}
			dirs = append(dirs, utils.Pt{C: v1, R: v2})
		}
	}

	return &solution{
		ans:    0,
		dirs:   dirs,
		target: "XMAS",
		StringGrid: utils.StringGrid{
			NCol:  len(lines[0]),
			NRow:  len(lines),
			Array: arr,
		},
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
