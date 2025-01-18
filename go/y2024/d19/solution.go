package main

import (
	"adventofcode/utils"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func (s *solution) run1() {
	for _, design := range s.desighs {
		s.memo = make(map[int]bool)
		if s.dfs3(0, design) {
			s.ans++
		}
	}
}

func (s *solution) dfs3(start int, design string) bool {
	if start == len(design) {
		return true
	}
	if val, ok := s.memo[start]; ok {
		return val
	}
	res := false
	for patternL := range s.itemLen {
		end := start + patternL
		if end > len(design) {
			continue
		}
		curpart := design[start:end]
		if s.patterns[curpart] {
			res = s.dfs(end, design)
			if res {
				break
			}
		}
	}
	s.memo[start] = res
	return res
}

func (s *solution) dfs(start int, design string) bool {
	if start == len(design) {
		return true
	}
	if val, ok := s.memo[start]; ok {
		return val
	}
	res := false
	for pattern := range s.patterns {
		end := start + len(pattern)
		if end > len(design) {
			continue
		}
		curpart := design[start:end]
		if s.patterns[curpart] {
			res = s.dfs(end, design)
			if res {
				break
			}
		}
	}
	s.memo[start] = res
	return res
}

func (s *solution) dfs2(start int, design string) int {
	if start == len(design) {
		return 1
	}
	if val, ok := s.memo2[start]; ok {
		return val
	}
	res := 0
	for patternL := range s.itemLen {
		end := start + patternL
		if end > len(design) {
			continue
		}
		curpart := design[start:end]
		if s.patterns[curpart] {
			res += s.dfs2(end, design)
		}
	}
	s.memo2[start] = res
	return res
}

func (s *solution) run2() {
	for _, design := range s.desighs {
		s.memo2 = make(map[int]int)
		s.ans += s.dfs2(0, design)
	}
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

	designs := []string{}
	itemLen := map[int]bool{}
	patterns := map[string]bool{}

	for i, line := range lines {
		if i == 0 {
			p := strings.Split(line, ", ")
			for _, v := range p {
				itemLen[len(v)] = true
				patterns[v] = true
			}
		}
		if i == 1 {
			continue
		}
		designs = append(designs, line)
	}
	return &solution{
		ans:      0,
		desighs:  designs,
		itemLen:  itemLen,
		patterns: patterns,
	}
}

type solution struct {
	ans           int
	desighs       []string
	itemLen, memo map[int]bool
	patterns      map[string]bool
	memo2         map[int]int
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
