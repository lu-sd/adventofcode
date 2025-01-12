package main

import (
	"adventofcode/utils"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

func (s *solution) run1() {
	for _, v := range s.nums {
		s.ans += s.dfs(v, 0, 75)
	}
}

func (s *solution) dfs(stone, cur, target int) int {
	if cur == target {
		return 1
	}
	id := fmt.Sprintf("%d,%d", stone, cur)
	if v, ok := s.memo[id]; ok {
		return v
	}
	newStones := s.blink(stone)
	ans := 0
	for _, v := range newStones {
		ans += s.dfs(v, cur+1, target)
	}
	s.memo[id] = ans
	return ans
}

func (s *solution) blink(stone int) (newStones []int) {
	if stone == 0 {
		newStones = append(newStones, 1)
	} else if len(strconv.Itoa(stone))%2 == 0 {
		numStr := strconv.Itoa(stone)
		half := len(numStr) / 2
		left, _ := strconv.Atoi(numStr[:half])
		right, _ := strconv.Atoi(numStr[half:])
		newStones = append(newStones, left, right)
	} else {
		newStones = append(newStones, stone*2024)
	}
	return
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

	nums := utils.IntsFromString(lines[0])
	return &solution{
		ans:  0,
		nums: nums,
		memo: map[string]int{},
	}
}

type solution struct {
	ans  int
	nums []int
	memo map[string]int
}

func part1(r io.Reader) int {
	s := buildSolution(r)
	s.run1()
	return s.res1()
}

func part2(r io.Reader) int {
	s := buildSolution(r)
	s.run1()
	return s.res1()
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
