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

func buildSolution(r io.Reader) *solution {
	lines, _ := h.LinesFromReader(r)
	return &solution{
		input: lines,
		ans1:  0,
		ans2:  0,
	}
}

func parseline(s string) (int, int) {
	parts := strings.FieldsFunc(s, func(r rune) bool {
		return r == ':' || r == '|'
	})

	lookUp := make(map[string]bool)
	var num int
	for win := range strings.FieldsSeq(parts[1]) {
		lookUp[win] = true
	}
	for item := range strings.FieldsSeq(parts[2]) {
		if ok := lookUp[item]; ok {
			num++
		}
	}
	if num > 0 {
		return num, 1 << (num - 1)
	}
	return num, 0
}

func (s *solution) run1() {
	for _, line := range s.input {
		_, score := parseline(line)
		s.ans1 += score
	}
}

func (s *solution) run2() {
	record := map[int]int{}
	fmt.Println("start m", record)
	for i, line := range s.input {
		count, _ := parseline(line)
		// update  based on previous record
		for j := range count {
			copyC := i + 1 + j
			// get a i+1 to copyC mapping
			// each of previous i+1 need trigger a +1 operation
			fmt.Print("update:", copyC, "  ")
			record[copyC] += record[i]
		}
		fmt.Println(record)
		// update current line observe
		for j := range count + 1 {
			copyC := i + 1 + j
			record[copyC] += 1
		}
		// record[i] += 1
	}
	for _, v := range record {
		s.ans2 += v
	}
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
