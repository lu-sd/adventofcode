package main

import (
	"adventofcode/utils"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"sort"
	"strings"
	"time"
)

type solution struct {
	ans        int
	nodes      []string
	nodemaps   map[string]struct{}
	connects   map[string]map[string]bool
	path       []string
	pathUnique map[string]bool
	result     []string
}

func (s *solution) run1() {
	s.dfs1(0)
}

func (s *solution) dfs1(idx int) {
	if len(s.path) == 3 {
		for _, v := range s.path {
			if v[0] == 't' {
				s.ans++
				break
			}
		}
		return
	}

	for i := idx; i < len(s.nodes); i++ {
		item := s.nodes[i]
		if s.pathUnique[item] {
			continue
		}
		if !s.all_conncted(item) {
			continue
		}
		s.pathUnique[item] = true
		s.path = append(s.path, item)
		s.dfs1(i)
		s.path = s.path[:len(s.path)-1]
		s.pathUnique[item] = false
	}
}

func (s *solution) dfs2(from string, edges map[string]bool) {
	if len(s.result) < len(s.path) {
		s.result = slices.Clone(s.path)
		fmt.Println(s.result)
	}

	for to := range edges {
		if to > from {
			continue
		}
		if s.pathUnique[to] {
			continue
		}
		if !s.all_conncted(to) {
			continue
		}
		s.pathUnique[to] = true
		s.path = append(s.path, to)
		nextEdges := s.connects[to]
		s.dfs2(to, nextEdges)
		s.path = s.path[:len(s.path)-1]
		s.pathUnique[to] = false
	}
}

func (s *solution) all_conncted(to string) bool {
	for _, from := range s.path {
		if !s.connects[from][to] {
			return false
		}
	}
	return true
}

func (s *solution) run2() {
	for from := range s.connects {
		s.path = append(s.path, from)
		s.dfs2(from, s.connects[from])
		s.path = s.path[:len(s.path)-1]
	}
}

func (s *solution) res() int {
	return s.ans
}

func buildSolution(r io.Reader) *solution {
	lines, err := utils.LinesFromReader(r)
	if err != nil {
		log.Fatalf("could not read input: %v %v", lines, err)
	}
	nodes := map[string]struct{}{}
	connects := map[string]map[string]bool{}

	for _, line := range lines {
		fromTo := strings.Split(line, "-")
		for i := range 2 {
			from, to := fromTo[i], fromTo[1-i]
			edges := connects[from]
			if edges == nil {
				edges = map[string]bool{}
				connects[from] = edges
			}
			connects[from][to] = true
			nodes[from] = struct{}{}
			nodes[to] = struct{}{}
		}
	}
	var lists []string
	for k := range nodes {
		lists = append(lists, k)
	}
	sort.Strings(lists)

	return &solution{
		nodes:      lists,
		nodemaps:   nodes,
		pathUnique: map[string]bool{},
		connects:   connects,
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
