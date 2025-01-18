package main

import (
	"adventofcode/utils"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type pos struct {
	utils.Pt
	step int
}

func (s *solution) run1() {
	s.bfs(utils.Pt{R: 0, C: 0}, 0)
}

func (s *solution) bfs(start utils.Pt, st int) {
	queue := []pos{{Pt: start, step: st}}
	s.Set(start, true)
	for len(queue) > 0 {
		n := len(queue)
		for n > 0 {
			l := queue[0]
			queue = queue[1:]
			if l.C == s.NCol-1 && l.R == s.NRow-1 {
				s.ans = l.step
				return
			}
			for _, dir := range utils.Dir4 {
				nextP := l.PMove(dir)
				if !s.IsInside(nextP) || s.Get(nextP) {
					continue
				}
				s.Set(nextP, true)
				queue = append(queue, pos{Pt: nextP, step: l.step + 1})
			}
			n--
		}
	}
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
	nrow := 71
	ncol := 71
	grid := make([][]bool, nrow)
	for i := range nrow {
		grid[i] = make([]bool, ncol)
	}

	for i, line := range lines {
		nums := utils.IntsFromString(line)
		grid[nums[1]][nums[0]] = true
		if i == 1023 {
			break
		}
	}
	return &solution{
		ans: 0,
		Grid: utils.Grid[bool]{
			NRow:  nrow,
			NCol:  ncol,
			Array: grid,
		},
	}
}

type solution struct {
	ans int
	utils.Grid[bool]
}

func part1(r io.Reader) int {
	s := buildSolution(r)
	s.run1()
	return s.res1()
}

func part2(r io.Reader) int {
	lines, err := utils.LinesFromReader(r)
	if err != nil {
		log.Fatalf("could not read input: %v %v", lines, err)
	}
	nrow := 71
	ncol := 71
	for checked, res := range lines {
		if checked <= 1023 {
			continue
		}

		grid := make([][]bool, nrow)
		for i := range nrow {
			grid[i] = make([]bool, ncol)
		}

		for i, line := range lines {
			nums := utils.IntsFromString(line)
			grid[nums[1]][nums[0]] = true
			if i == checked {
				break
			}
		}
		s := &solution{
			ans: 0,
			Grid: utils.Grid[bool]{
				NRow:  nrow,
				NCol:  ncol,
				Array: grid,
			},
		}
		s.run1()
		if s.ans == 0 {
			fmt.Println(checked, res)
			break
		}
	}
	return 0
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
