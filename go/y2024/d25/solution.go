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
	for _, key := range s.pairs['.'] {
		for _, lock := range s.pairs['#'] {
			if s.valid(key, lock) {
				s.ans++
			}
		}
	}
}

func (s *solution) valid(a, b []int) bool {
	for i, v := range a {
		if v+b[i] > 5 {
			return false
		}
	}
	return true
}

func (s *solution) run2() {
}

func (s *solution) res() int {
	return s.ans
}

func Lockvalue(runeG [][]rune) []int {
	nrow, ncol := len(runeG), len(runeG[0])
	res := make([]int, ncol)
	for j := range ncol {
		count := -1
		for i := range nrow {
			if runeG[i][j] == '#' {
				count++
			}
		}
		res[j] = count
	}
	return res
}

func buildSolution(r io.Reader) *solution {
	lines, err := utils.LinesFromReader(r)
	if err != nil {
		log.Fatalf("could not read input: %v %v", lines, err)
	}
	pairs := map[rune][][]int{}
	hint := '0'
	var runeG [][]rune
	for _, line := range lines {
		if len(line) == 0 {
			if len(runeG) > 0 {
				pairs[hint] = append(pairs[hint], Lockvalue(runeG))
			}
			hint = '0'
			runeG = runeG[:0]
			continue
		}
		if line[0] == '#' && hint == '0' {
			hint = '#'
		}
		if line[0] == '.' && hint == '0' {
			hint = '.'
		}
		runeG = append(runeG, []rune(line))
	}
	if len(runeG) > 0 {
		pairs[hint] = append(pairs[hint], Lockvalue(runeG))
	}

	// fmt.Printf("%#v", pairs)

	return &solution{
		pairs: pairs,
		ans:   0,
	}
}

type solution struct {
	pairs map[rune][][]int
	ans   int
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
