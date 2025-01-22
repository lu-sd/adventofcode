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
	for _, keys := range s.pairs['#'] {
		for _, locks := range s.pairs['.'] {
			if s.isValid(keys, locks) {
				s.ans++
			}
		}
	}
}

func (s *solution) isValid(key, lock []int) bool {
	for _, k := range key {
		for _, l := range lock {
			if k+l > 5 {
				return false
			}
		}
	}
	return true
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
	grid := utils.StringGrid{
		Array: []string{},
	}

	hint := '0'
	pairs := make(map[rune][][]int)

	getHeight := func(g utils.StringGrid) []int {
		nrow := len(g.Array)
		ncol := len(g.Array[0])
		height := make([]int, ncol)
		for c := 0; c < ncol; c++ {
			count := -1
			for r := 0; r < nrow; r++ {
				if g.GetRune(utils.Pt{R: r, C: c}) == '#' {
					count++
				}
			}
			height[c] = count
		}
		return height
	}

	for _, line := range lines {
		if hint == '0' {
			hint = rune(line[0])
		}

		if len(line) == 0 {
			pairs[hint] = append(pairs[hint], getHeight(grid))
			grid.Array = []string{}
			// grid.Array = grid.Array[:0]
			hint = '0'
			continue
		}
		grid.Array = append(grid.Array, line)

	}
	pairs[hint] = append(pairs[hint], getHeight(grid))
	fmt.Println(pairs)

	return &solution{
		ans:   0,
		pairs: pairs,
	}
}

type solution struct {
	ans   int
	pairs map[rune][][]int
}

func part1(r io.Reader) int {
	s := buildSolution(r)
	s.run1()
	return s.res1()
}

func part2(r io.Reader) int {
	// s := buildSolution(r)
	// s.run2()
	// return s.res2()
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
