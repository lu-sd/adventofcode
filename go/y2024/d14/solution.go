package main

import (
	"adventofcode/utils"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func (s *solution) run1() {
	for _, r := range s.robots {
		r.C = (r.C + r.vx*second) % ncol
		r.R = (r.R + r.vy*second) % nrow
		if r.C < 0 {
			r.C += ncol
		}
		if r.R < 0 {
			r.R += nrow
		}
	}
}

func (s *solution) run2() {
}

func (s *solution) res1() int {
	quadrant := [4]int{}
	for _, r := range s.robots {
		if r.C < ncol/2 && r.R < nrow/2 {
			quadrant[0]++
		} else if r.C < ncol/2 && r.R > nrow/2 {
			quadrant[1]++
		} else if r.C > ncol/2 && r.R > nrow/2 {
			quadrant[2]++
		} else if r.C > ncol/2 && r.R < nrow/2 {
			quadrant[3]++
		}
	}
	s.ans = 1
	for i := range 4 {
		s.ans *= quadrant[i]
	}
	return s.ans
}

func (s *solution) res2() int {
	return s.ans
}

func buildSolution(r io.Reader) *solution {
	lines, err := utils.LinesFromReader(r)
	if err != nil {
		log.Fatalf("could not read input: %v %v", lines, err)
	}
	robots := []robot{}

	for _, line := range lines {
		nums := utils.IntsFromString(line)
		rb := robot{
			utils.Pt{C: nums[0], R: nums[1]},
			nums[2],
			nums[3],
		}
		robots = append(robots, rb)
	}

	return &solution{
		ans:    0,
		robots: robots,
	}
}

type robot struct {
	utils.Pt
	vx, vy int
}

const (
	nrow   = 103
	ncol   = 101
	second = 100
)

type solution struct {
	ans    int
	robots []robot
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
	Input, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("fail open input.txt %v", err)
	}
	start := time.Now()
	result := part1(Input)
	elapsed := time.Since(start)
	fmt.Printf("p1 res 🙆-> %d (Time taken: %s)\n", result, elapsed)
	start = time.Now()
	result = part2(Input)
	elapsed = time.Since(start)
	fmt.Printf("p2 res 🙆-> %d (Time taken: %s)\n", result, elapsed)
}
