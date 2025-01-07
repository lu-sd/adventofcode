package main

import (
	"adventofcode/utils"
	"fmt"
	"io"
	"log"
	"os"
)

type (
	anti     map[rune][]utils.Pt
	antiNode map[utils.Pt]struct{}
)

type solution struct {
	ans, nRow, nCol int
	anti            anti
	antiNodes       antiNode
	dirs            [2]int
}

func (s *solution) isInside(pt utils.Pt) bool {
	return pt.C < s.nCol && pt.C >= 0 && pt.R < s.nRow && pt.R >= 0
}

func (s *solution) run1() {
	for _, pts := range s.anti {
		for i := range pts {
			for j := i + 1; j < len(pts); j++ {
				ptParis := [2]utils.Pt{
					pts[i], pts[j],
				}
				s.findNode(ptParis)
			}
		}
	}
}

func (s *solution) findNode(ptParis [2]utils.Pt) {
	dx, dy := ptParis[0].Dist(ptParis[1])
	for idx, pt := range ptParis {
		nextPt := pt.Move(dx*s.dirs[idx], dy*s.dirs[idx])
		if s.isInside(nextPt) {
			s.antiNodes[nextPt] = struct{}{}
		}
	}
}

func (s *solution) findNode2(ptParis [2]utils.Pt) {
	dx, dy := ptParis[0].Dist(ptParis[1])
	for idx, pt := range ptParis {
		s.antiNodes[pt] = struct{}{}
		step := 1
		for {
			nextPt := pt.Move(dx*s.dirs[idx]*step, dy*s.dirs[idx]*step)
			if s.isInside(nextPt) {
				s.antiNodes[nextPt] = struct{}{}
				step++
			} else {
				break
			}
		}
	}
}

func (s *solution) run2() {
	for _, pts := range s.anti {
		for i := range pts {
			for j := i + 1; j < len(pts); j++ {
				ptParis := [2]utils.Pt{
					pts[i], pts[j],
				}
				s.findNode2(ptParis)
			}
		}
	}
}

func (s *solution) res() int {
	return len(s.antiNodes)
}

func buildSolution(r io.Reader) *solution {
	lines, err := utils.LinesFromReader(r)
	if err != nil {
		log.Fatalf("could not read input: %v %v", lines, err)
	}
	anti := anti{}
	for i, line := range lines {
		for j, rune := range line {
			if rune == '.' {
				continue
			}
			anti[rune] = append(anti[rune], utils.Pt{C: j, R: i})
		}
	}

	return &solution{
		ans:       0,
		anti:      anti,
		antiNodes: antiNode{},
		nRow:      len(lines),
		nCol:      len(lines[0]),
		dirs:      [2]int{1, -1},
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
