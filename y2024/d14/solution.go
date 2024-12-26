package main

import (
	"adventofcode/utils"
	"fmt"
	"io"
	"log"
	"os"
)

type robot struct {
	p, v utils.Pt
}

type solution struct {
	ans    int
	robots []robot
}

const (
	width, height = 101, 103
	seconds       = 100
)

func (s *solution) run1() {
	for i := range s.robots {
		s.robots[i].p.C = (s.robots[i].p.C + s.robots[i].v.C*seconds) % width
		s.robots[i].p.R = (s.robots[i].p.R + s.robots[i].v.R*seconds) % height
		if s.robots[i].p.C < 0 {
			s.robots[i].p.C += width
		}
		if s.robots[i].p.R < 0 {
			s.robots[i].p.R += height
		}
	}
}

func (s *solution) run2() {
}

func (s *solution) res() int {
	quadrants := [4]int{}
	for _, r := range s.robots {
		if r.p.C < width/2 && r.p.R < height/2 {
			quadrants[0]++
		} else if r.p.C < width/2 && r.p.R > height/2 {
			quadrants[1]++
		} else if r.p.C > width/2 && r.p.R > height/2 {
			quadrants[2]++
		} else if r.p.C > width/2 && r.p.R < height/2 {
			quadrants[3]++
		}
	}
	s.ans = 1
	for i := range 4 {
		s.ans *= quadrants[i]
	}
	// s.ans = quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
	return s.ans
}

func buildSolution(r io.Reader) *solution {
	lines, err := utils.LinesFromReader(r)
	if err != nil {
		log.Fatalf("could not read input: %v %v", lines, err)
	}
	robots := []robot{}

	for _, line := range lines {
		ints := utils.IntsFromString(line)
		rb := robot{
			utils.Pt{C: ints[0], R: ints[1]},
			utils.Pt{C: ints[2], R: ints[3]},
		}
		robots = append(robots, rb)
	}
	return &solution{
		ans:    0,
		robots: robots,
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
	Input, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("fail open test1.txt %v", err)
	}
	defer Input.Close()
	fmt.Println("p1 res 🙆-> ", part1(Input))
	fmt.Println("p2 res 🙆-> ", part2(Input))
}
