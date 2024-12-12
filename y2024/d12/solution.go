package main

import (
	"adventofcode/utils"
	"fmt"
	"io"
	"log"
	"os"
)

type (
	headPt   utils.Pt
	solution struct {
		ans   int
		lands map[headPt]map[utils.Pt]rune
		seen  map[utils.Pt]bool
		utils.Grid
	}
)

func (s *solution) run1() {
	for i, line := range s.Array {
		for j, flower := range line {
			cur := utils.Pt{X: i, Y: j}
			if s.seen[cur] {
				continue
			}
			s.dfs1(cur, flower, headPt(cur))
		}
	}
}

func (s *solution) res() int {
	for _, land := range s.lands {
		area, perimeter := len(land), 0
		for pt, flower := range land {
			for _, dir := range utils.Dir4 {
				nexPt := pt.Move(dir.X, dir.Y)
				if !s.IsInside(nexPt) || s.PRune(nexPt) != flower {
					perimeter++
				}
			}
		}
		s.ans += area * perimeter
	}
	return s.ans
}

func (s *solution) dfs1(curP utils.Pt, flower rune, header headPt) {
	if !s.IsInside(curP) || s.PRune(curP) != flower || s.seen[curP] {
		return
	}
	s.seen[curP] = true
	if s.lands[header] == nil {
		s.lands[header] = map[utils.Pt]rune{}
	}
	s.lands[header][curP] = flower
	for _, dir := range utils.Dir4 {
		nextP := curP.Move(dir.X, dir.Y)
		s.dfs1(nextP, flower, header)
	}
}

func (s *solution) run2() {
}

func buildSolution(r io.Reader) *solution {
	lines, err := utils.LinesFromReader(r)
	if err != nil {
		log.Fatalf("could not read input: %v %v", lines, err)
	}
	nrow, ncol := len(lines), len(lines[0])

	return &solution{
		ans:   0,
		seen:  map[utils.Pt]bool{},
		lands: map[headPt]map[utils.Pt]rune{},
		Grid: utils.Grid{
			NRow:  nrow,
			NCol:  ncol,
			Array: lines,
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
	s.run2()
	return s.res()
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
