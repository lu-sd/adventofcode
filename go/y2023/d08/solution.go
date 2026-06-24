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

type Node struct {
	left, right string
}

func buildSolution(r io.Reader) *solution {
	lines, _ := h.LinesFromReader(r)
	return &solution{
		input: lines,
		ans1:  0,
		ans2:  0,
	}
}

func navigate(lines []string) (ins string, nodes map[string]Node) {
	nodes = make(map[string]Node)
	for i, line := range lines {
		if i == 0 {
			ins = line
			continue
		}
		if line == "" {
			continue
		}
		var key, left, right string
		fmt.Sscanf(line, "%3s = (%3s, %3s)", &key, &left, &right)
		nodes[key] = Node{
			left:  left,
			right: right,
		}
	}
	return
}

//	func parseLine(l string) (node map[string]Node) {
//		node = make(map[string]Node)
//		key, value, _ := strings.Cut(l, "=")
//		value = strings.Trim(value, "()")
//		left, right, _ := strings.Cut(value, ",")
//		node[key] = Node{left, right}
//		return node
//	}
func (s *solution) run1() {
	ins, nodes := navigate(s.input)
	start := "AAA"
	end := "ZZZ"
	current := start

	for current != end {
		for _, c := range ins {
			node := nodes[current]
			if c == 'L' {
				current = node.left
			} else {
				current = node.right
			}
			s.ans1++
			if current == end {
				break
			}
		}
	}
}

func (s *solution) run2() {
	ins, nodes := navigate(s.input)

	var currents []string
	for key := range nodes {
		if strings.HasSuffix(key, "A") {
			currents = append(currents, key)
		}
	}

	for {
		for _, c := range ins {
			for i, curr := range currents {
				node := nodes[curr]

				if c == 'L' {
					currents[i] = node.left
				} else {
					currents[i] = node.right
				}
			}

			s.ans2++

			allEndWithZ := true
			for _, cur := range currents {
				if !strings.HasSuffix(cur, "Z") {
					allEndWithZ = false
					break
				}
			}

			if allEndWithZ {
				return
			}
		}
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
