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
		ans int
		// lands map[headPt]map[utils.Pt]rune
		lands map[headPt][]utils.Pt
		seen  map[utils.Pt]bool
		utils.Grid[rune]
	}
)

func (s *solution) run1() {
	for i, line := range s.Array {
		for j, flower := range line {
			cur := utils.Pt{C: i, R: j}
			if s.seen[cur] {
				continue
			}
			s.dfs1(cur, flower, headPt(cur))
		}
	}
}

func (s *solution) res() int {
	for header, land := range s.lands {
		area, perimeter := len(land), 0
		flower := s.Get(utils.Pt(header))
		for _, pt := range land {
			for _, dir := range utils.Dir4 {
				nexPt := pt.Move(dir.C, dir.R)
				if !s.IsInside(nexPt) || s.Get(nexPt) != flower {
					perimeter++
				}
			}
		}
		s.ans += area * perimeter
	}
	return s.ans
}

func (s *solution) res2() int {
	for header, land := range s.lands {
		area, sides := len(land), 0
		flower := s.Get(utils.Pt(header))
		for _, pt := range land {
			boolSlice := make([]bool, 4)
			// check whether neigher direc with same bool
			for i, dir := range utils.Dir4 {
				nexPt := pt.Move(dir.C, dir.R)
				if !s.IsInside(nexPt) || s.Get(nexPt) != flower {
					boolSlice[i] = true
				}
			}
			for i, v := range boolSlice {
				neighb := (i + 1) % 4
				// out conner, both neighb directioin can not go
				if v && boolSlice[neighb] {
					sides++
				}
				// for insider conner, both can go
				if !v && !boolSlice[neighb] {
					pt1, pt2 := utils.Dir4[i], utils.Dir4[neighb]
					anglePt := pt.Move(pt1.C+pt2.C, pt1.R+pt2.R)
					if s.Get(anglePt) != flower {
						sides++
					}
				}
			}
		}
		// fmt.Printf("flower %c area %d slides %d \n", flower, area, sides)
		s.ans += area * sides
	}
	return s.ans
}

func (s *solution) dfs1(curP utils.Pt, flower rune, header headPt) {
	if !s.IsInside(curP) || s.Get(curP) != flower || s.seen[curP] {
		return
	}
	s.seen[curP] = true
	// if s.lands[header] == nil {
	// 	s.lands[header] = map[utils.Pt]rune{}
	// }

	s.lands[header] = append(s.lands[header], curP)
	for _, dir := range utils.Dir4 {
		nextP := curP.Move(dir.C, dir.R)
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
	rline := [][]rune{}
	for _, v := range lines {
		rline = append(rline, []rune(v))
	}

	return &solution{
		ans:   0,
		seen:  map[utils.Pt]bool{},
		lands: map[headPt][]utils.Pt{},
		Grid: utils.Grid[rune]{
			NRow:  nrow,
			NCol:  ncol,
			Array: rline,
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
	s.run1()
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
