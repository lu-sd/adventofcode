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
	s.bfs(s.start, utils.Dir4[3])
}

func (s *solution) bfs(curP, dir utils.Pt) {
	queue := []sp{}
	queue = append(queue, sp{curP, dir, 0})
	// queue = []sp{{curP,dir,0}}
	s.Seen[curP] = true
	for len(queue) > 0 {
		sz := len(queue)
		for sz > 0 {
			top := queue[0]
			queue = queue[1:]

			if s.GetRune(top.pos) == 'E' {
				if s.ans == 0 || top.score < s.ans {
					s.ans = top.score
				}
			}

			for _, dir := range utils.Dir4 {
				nextP := top.pos.PMove(dir)
				if s.GetRune(nextP) == '#' {
					continue
				}
				turn := 0
				if top.dir != dir {
					turn = 1
				}
				nextS := top.score + 1 + 1000*turn
				if !s.Seen[nextP] {
					s.Seen[nextP] = true
					queue = append(queue, sp{nextP, dir, nextS})
				}

			}
			sz--
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

	return &solution{
		ans:   0,
		start: utils.Pt{C: 1, R: len(lines) - 2},
		StringGrid: utils.StringGrid{
			NCol:  len(lines[0]),
			NRow:  len(lines),
			Array: lines,
		},
		Seen: utils.Seen{},
	}
}

type sp struct {
	pos, dir utils.Pt
	score    int
}

type solution struct {
	ans int
	utils.StringGrid
	utils.Seen
	start utils.Pt
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
