package main

import (
	"adventofcode/utils"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var dirs = map[rune]utils.Pt{
	'^': {C: 0, R: -1},
	'>': {C: 1, R: 0},
	'<': {C: -1, R: 0},
	'v': {C: 0, R: 1},
}

func (s *solution) run1() {
	// s.Set(s.robot, '.')
	s.Array[s.robot.R][s.robot.C] = '.'
	for _, r := range s.instrucions {
		dir := dirs[r]
		nextP := s.robot.PMove(dir)
		if s.Get(nextP) == '#' {
			continue
		}
		if s.Get(nextP) == '.' {
			s.robot = nextP
			continue
		}
		if spot, find := s.Find(s.robot, dir, '.', '#'); find {
			s.Swap(spot, nextP)
			s.robot = nextP
		}
	}
}

func (s *solution) run2() {
}

func (s *solution) res1() int {
	for i, line := range s.Array {
		for j, r := range line {
			if r == 'O' {
				s.ans += 100*i + j
			}
		}
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

	nrow := 0
	ncol := len(lines[0])
	first := true
	robot := utils.Pt{}
	array := make([][]rune, 0)
	instrucions := make([]rune, 0)

	for i, line := range lines {
		if len(line) == 0 {
			first = false
			nrow = i
			continue
		}

		if first {
			for j, rune := range line {
				if rune == '@' {
					robot = utils.Pt{C: j, R: i}
				}
			}
			array = append(array, []rune(line))
		} else {
			instrucions = append(instrucions, []rune(line)...)
		}
	}
	return &solution{
		ans: 0,
		Grid: utils.Grid[rune]{
			NRow:  nrow,
			NCol:  ncol,
			Array: array,
		},
		robot:       robot,
		instrucions: instrucions,
	}
}

type solution struct {
	ans int
	utils.Grid[rune]
	instrucions []rune
	robot       utils.Pt
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
