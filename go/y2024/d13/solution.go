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
	for _, m := range s.machines {
		s.ans += s.lowest(m.a, m.b, m.target)
	}
}

func (s *solution) lowest(a, b, target utils.Pt) int {
	lowest, found := 0, false
	for aNum := range s.limit {
		for bNum := range s.limit {
			if s.found(aNum, bNum, a, b, target) {
				cur := aNum*3 + bNum
				if !found {
					lowest = cur
				} else {
					lowest = min(lowest, cur)
				}
			}
		}
	}
	return lowest
}

func (s *solution) found(aNum, bNum int, a, b, target utils.Pt) bool {
	return (aNum*a.C+bNum*b.C) == target.C && (aNum*a.R+bNum*b.R) == target.R
}

func (s *solution) run2() {
}

func (s *solution) res1() int {
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

	machines := make([]machine, 0)
	pt := [3]utils.Pt{}
	for i, line := range lines {
		nums := utils.IntsFromString(line)
		if i%4 == 3 {
			machines = append(machines, machine{pt[0], pt[1], pt[2]})
			continue
		}
		pt[i%4].C = nums[0]
		pt[i%4].R = nums[1]
	}
	machines = append(machines, machine{pt[0], pt[1], pt[2]})

	return &solution{
		machines: machines,
		limit:    100,
		ans:      0,
	}
}

type machine struct {
	a, b, target utils.Pt
}
type solution struct {
	ans      int
	machines []machine
	limit    int
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
