package main

import (
	"adventofcode/h"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

type solution struct {
	input      []string
	ans1, ans2 int
}
type R struct {
	dst, src, dis int
}

func buildSolution(r io.Reader) *solution {
	lines, _ := h.LinesFromReader(r)
	return &solution{
		input: lines,
		ans1:  0,
		ans2:  0,
	}
}

func parse(s []string) (seeds []int, relations [][]R) {
	for i := 0; i < len(s); i++ {

		if i == 0 {
			_, seedStr, _ := strings.Cut(s[i], ":")
			seeds = h.IntsFromString(seedStr)
			continue
		}

		if s[i] == "" {
			continue
		}

		var rel []R

		for ; i < len(s) && s[i] != ""; i++ {
			v := h.IntsFromString(s[i])
			if len(v) != 0 {
				rel = append(rel, R{v[0], v[1], v[2]})
			}
		}
		relations = append(relations, rel)
	}
	return
}

func findDst(seed int, relation []R) (res int) {
	for _, v := range relation {
		if seed >= v.src && seed < v.src+v.dis {
			res = seed - v.src + v.dst
			return res
		}
	}
	return seed
}

func (s *solution) run1() {
	s.ans1 = 1<<32 - 1
	seeds, relations := parse(s.input)
	for _, seed := range seeds {
		for _, relation := range relations {
			seed = findDst(seed, relation)
		}
		s.ans1 = min(seed, s.ans1)
	}
}

func (s *solution) run2() {
	s.ans2 = 1<<32 - 1
	seeds, relations := parse(s.input)
	seeds = dealSeeds(seeds)
	for _, seed := range seeds {
		for _, relation := range relations {
			seed = findDst(seed, relation)
		}
		s.ans2 = min(seed, s.ans2)
	}
}

func dealSeeds(s []int) (seeds []int) {
	for i := 0; i < len(s); i = i + 2 {
		for j := 0; j < s[i+1]; j++ {
			newSeed := s[i] + j
			seeds = append(seeds, newSeed)
		}
	}
	return seeds
}

func (s *solution) res1() int {
	return s.ans1
}

func (s *solution) res2() int {
	return s.ans2
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
