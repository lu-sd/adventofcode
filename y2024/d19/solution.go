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
	patternList map[string]bool
	designs     []string
}

func (s *solution) dfs(start int, design string) bool {
	if start == len(design) {
		return true
	}
	for i := start; i < len(design); i++ {
		prefix := design[start : i+1]
		if s.patternList[prefix] && s.dfs(i+1, design) {
			return true
		}
	}
	return false
}

func (s *solution) run1() {
	fmt.Printf("%#v \n", s)
	// fmt.Println(s)
	for _, design := range s.designs {
		if s.dfs(0, design) {
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
	pattern := strings.Split(lines[0], ", ")
	patternList := make(map[string]bool)
	for _, w := range pattern {
		patternList[w] = true
	}
	design := []string{}
	for i, line := range lines {
		if i <= 1 {
			continue
		}
		design = append(design, line)
	}

	return &solution{
		ans:         0,
		patternList: patternList,
		designs:     design,
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
