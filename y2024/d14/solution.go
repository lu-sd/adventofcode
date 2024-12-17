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
		s.robots[i].p.X = (s.robots[i].p.X + s.robots[i].v.X*seconds) % width
		s.robots[i].p.Y = (s.robots[i].p.Y + s.robots[i].v.Y*seconds) % height
		if s.robots[i].p.X < 0 {
			s.robots[i].p.X += width
		}
		if s.robots[i].p.Y < 0 {
			s.robots[i].p.Y += height
		}
	}
}

func (s *solution) run2() {
}

func (s *solution) res() int {
	quadrants := [4]int{}
	for _, r := range s.robots {
		if r.p.X < width/2 && r.p.Y < height/2 {
			quadrants[0]++
		} else if r.p.X < width/2 && r.p.Y > height/2 {
			quadrants[1]++
		} else if r.p.X > width/2 && r.p.Y > height/2 {
			quadrants[2]++
		} else if r.p.X > width/2 && r.p.Y < height/2 {
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
			utils.Pt{X: ints[0], Y: ints[1]},
			utils.Pt{X: ints[2], Y: ints[3]},
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
