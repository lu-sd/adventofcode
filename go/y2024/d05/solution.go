package main

import (
	"adventofcode/utils"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
)

type (
	adj      map[int][]int
	solution struct {
		ans     int
		adjList adj
		manuals [][]int
	}
)

func (s *solution) run1() {
	for _, manual := range s.manuals {
		if s.isRightOrder(manual) {
			s.ans += s.middle(manual)
		}
	}
}

func (s *solution) isRightOrder(m []int) bool {
	seen := map[int]bool{}
	for _, num := range m {
		for _, after := range s.adjList[num] {
			if seen[after] {
				return false
			}
		}
		seen[num] = true
	}
	return true
}

func (s *solution) middle(m []int) int {
	return m[len(m)/2]
}

func (s *solution) run2() {
	for _, manual := range s.manuals {
		if !s.isRightOrder(manual) {
			s.ans += s.resorted(manual)
		}
	}
}

func (s *solution) buildList(task []int) adj {
	nAdjList := make(adj)
	for _, par := range task {
		dependcy := []int{}
		for _, num := range s.adjList[par] {
			if slices.Contains(task, num) {
				dependcy = append(dependcy, num)
			}
		}
		nAdjList[par] = dependcy
	}
	return nAdjList
}

func findIndegre(task []int, nAdjList adj) map[int]int {
	inDegree := make(map[int]int)
	for _, dependcy := range nAdjList {
		for _, num := range dependcy {
			inDegree[num]++
		}
	}
	return inDegree
}

func validAdj(m []int, indegree map[int]int, nAdjList adj) int {
	goodList := []int{}
	queue := []int{}
	for _, node := range m {
		if indegree[node] == 0 {
			queue = append(queue, node)
		}
	}

	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]
		goodList = append(goodList, head)
		for _, num := range nAdjList[head] {
			indegree[num]--
			if indegree[num] == 0 {
				queue = append(queue, num)
			}
		}
	}
	return goodList[len(goodList)/2]
}

func (s *solution) resorted(m []int) int {
	// buildList
	nAdjList := s.buildList(m)
	// findIndegre
	indegree := findIndegre(m, nAdjList)
	// validList
	res := validAdj(m, indegree, nAdjList)
	return res
}

func (s *solution) res() int {
	return s.ans
}

func buildSolution(r io.Reader) *solution {
	lines, err := utils.LinesFromReader(r)
	if err != nil {
		log.Fatalf("could not read input: %v %v", lines, err)
	}
	processAdj := true
	adjList := make(adj)
	manuals := make([][]int, 0)
	for _, line := range lines {
		nums := utils.IntsFromString(line)
		if len(line) == 0 {
			processAdj = false
			continue
		}
		if processAdj {
			adjList[nums[0]] = append(adjList[nums[0]], nums[1])
		} else {
			manuals = append(manuals, nums)
		}
	}
	return &solution{
		ans:     0,
		adjList: adjList,
		manuals: manuals,
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
