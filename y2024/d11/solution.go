package main

import (
	"adventofcode/utils"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type solution struct {
	stones []int
	ans    int
}

func (s *solution) bfs() []int {
	newStones := []int{}
	for _, stone := range s.stones {
		if stone == 0 {
			newStones = append(newStones, 1)
		} else if len(strconv.Itoa(stone))%2 == 0 {
			numstr := strconv.Itoa(stone)
			half := len(numstr) / 2
			leftN, _ := strconv.Atoi(numstr[:half])
			rightN, _ := strconv.Atoi(numstr[half:])
			newStones = append(newStones, leftN, rightN)
		} else {
			newStones = append(newStones, stone*2024)
		}
	}
	return newStones
}

func (s *solution) dfs(i, level, target int) int {
	if level == target {
		return 1
	}
	for _, stone := range s.stones {
		s.ans += s.dfs(stone, level+1, target)
	}
	return 0
}

func (s *solution) run1() {
	for range 25 {
		s.stones = s.bfs()
	}
}

func (s *solution) run2() {
	for _, stone := range s.stones {
		dfs(stone, 0, 75)
	}
}

func (s *solution) res() int {
	s.ans = len(s.stones)
	return s.ans
}

func buildSolution(r io.Reader) *solution {
	lines, err := utils.LinesFromReader(r)
	if err != nil {
		log.Fatalf("could not read input: %v %v", lines, err)
	}

	strNums := strings.Fields(lines[0])
	stones := []int{}

	for _, stone := range strNums {
		num, _ := strconv.Atoi(stone)
		stones = append(stones, num)
	}

	return &solution{
		ans:    0,
		stones: stones,
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
