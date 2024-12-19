package main

import (
	"adventofcode/utils"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type solution struct {
	ans         int
	patternList map[string]struct{}
	design      []string
}

func (s *solution) dfs(start int, w string) bool {
	if start == len(w) {
		return true
	}
	for i := start; i < len(w); i++ {
		prefix := w[start : i+1]
		if _, ok := s.patternList[prefix]; ok && s.dfs(i+1, w) {
			return true
		}
	}
	return false
}

func (s *solution) run1() {
	for _, word := range s.design {
		if s.dfs(0, word) {
			s.ans++
		}
	}
}

func (s *solution) run2() {
}

func (s *solution) res() int {
	return s.ans
}

func buildSolution(r io.Reader) *solution {
	lines, err := utils.LinesFromReader(r)
	if err != nil {
		log.Fatalf("could not read input: %v %v", lines, err)
	}
	pattern := strings.Split(lines[0], ",")
	patternList := make(map[string]struct{})
	for _, w := range pattern {
		patternList[w] = struct{}{}
	}
	design := []string{}
	for i, line := range lines {
		if i < 1 {
			continue
		}
		design = append(design, line)
	}

	return &solution{
		ans:         0,
		patternList: patternList,
		design:      design,
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
	arg := os.Args[1]
	fmt.Println("Running part", arg)
	switch arg {
	case "1":
		fmt.Println("p1 res 🙆-> ", part1(os.Stdin))
	case "2":
		fmt.Println("p2 res 🙆-> ", part2(os.Stdin))
	}
}
