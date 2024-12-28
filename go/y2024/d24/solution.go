package main

import (
	"adventofcode/utils"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func binaryToDecimal(binary []rune) int {
	res := 0
	for _, char := range binary {
		res *= 2
		res += int(char - '0')
	}
	// binary := strconv.FormatInt(int64(decimal), 2)
	return res
}

func (s *solution) run1() {
	s.buildGaph(s.graphlines)
	s.buildInDegree()
	sorted := s.topSort()
	// fmt.Printf("%#v \n", s.inDegree)
	// fmt.Printf("%#v \n", s.graph)
	// fmt.Printf("%#v \n", sorted)
	// fmt.Printf("%v \n", s.vertices)
	// for k, v := range s.vertices {
	// 	fmt.Printf("id %s %v \n", k, *v)
	// }
	// fmt.Println(s.inDegree)
	for _, id := range sorted {
		item := s.vertices[id]
		item.process(s.vertices)
	}
	// collect all z related id
	var res []string
	for _, id := range sorted {
		if id[0] == 'z' {
			res = append(res, id)
		}
	}
	sort.Strings(res)
	// fmt.Printf("%#v \n", res)
	binary := make([]rune, len(res))
	for i, id := range res {
		value := '0'
		if s.vertices[id].res {
			value = '1'
		}
		fmt.Println("id", id, string(value))
		binary[len(res)-i-1] = value
	}
	fmt.Println(string(binary))
	fmt.Println(binaryToDecimal(binary))
}

func (s *solution) run2() {
	s.buildGaph(s.graphlines)
	s.buildInDegree()
	sorted := s.topSort()
	// fmt.Printf("%#v \n", s.inDegree)
	// fmt.Printf("%#v \n", s.graph)
	// fmt.Printf("%#v \n", sorted)
	// fmt.Printf("%v \n", s.vertices)
	// for k, v := range s.vertices {
	// 	fmt.Printf("id %s %v \n", k, *v)
	// }
	// fmt.Println(s.inDegree)
	for _, id := range sorted {
		item := s.vertices[id]
		item.process(s.vertices)
	}
	// collect all z related id
	var res []string
	for _, id := range sorted {
		if id[0] == 'z' {
			res = append(res, id)
		}
	}
	sort.Strings(res)
	// fmt.Printf("%#v \n", res)
	binary := make([]rune, len(res))
	for i, id := range res {
		value := '0'
		if s.vertices[id].res {
			value = '1'
		}
		fmt.Println("id", id, string(value))
		binary[len(res)-i-1] = value
	}
	fmt.Println(string(binary))
	fmt.Println(binaryToDecimal(binary))
	z := strconv.FormatInt(int64(s.expectedZ), 2)

	fmt.Println("extected z:", s.expectedZ)
	fmt.Println("extected:", z)
}

func (s *solution) res() int {
	return s.ans
}

func (s *solution) buildInDegree() {
	inDegree := map[string]int{}
	for _, connects := range s.graph {
		for _, id := range connects {
			// add count for each small island
			inDegree[id]++ // dont care who point to this id, only count matter
			// x00 may miss, default is 0. which is expected
		}
	}
	s.inDegree = inDegree
}

func (s *solution) topSort() []string {
	var result, queue []string

	for id := range s.vertices {
		if s.inDegree[id] == 0 {
			queue = append(queue, id)
		}
	}
	for len(queue) > 0 {
		id := queue[0]
		queue = queue[1:]
		result = append(result, id)
		for _, child := range s.graph[id] {
			s.inDegree[child]--
			if s.inDegree[child] == 0 {
				queue = append(queue, child)
			}
		}
	}
	return result
}

func (s *solution) buildGaph(lines []string) {
	graph := map[string][]string{}
	for _, line := range lines {
		var a, op, b, id string
		fmt.Sscanf(line, "%s %s %s -> %s", &a, &op, &b, &id)
		// expand graph
		s.vertices[id] = &vertex{a: a, b: b, op: op, id: id}
		graph[a] = append(graph[a], id)
		graph[b] = append(graph[b], id)
	}
	s.graph = graph
}

func buildSolution(r io.Reader) *solution {
	lines, err := utils.LinesFromReader(r)
	if err != nil {
		log.Fatalf("could not read input: %v %v", lines, err)
	}
	// graph := map[string][]string{}
	// inDegree := map[string]int{}
	part1 := true
	var graphLines []string
	vertices := map[string]*vertex{} // normal only need string list
	var allX, allY string
	for _, line := range lines {
		if len(line) == 0 {
			part1 = false
			continue
		}
		if part1 {
			items := strings.Split(line, ": ")
			// fmt.Sscanf(line, "%s: %d", &id, &v)
			vertices[items[0]] = &vertex{res: items[1] == "1"}
			if items[0][0] == 'x' {
				allX = items[1] + allX
			}
			if items[0][0] == 'y' {
				allY = items[1] + allY
			}
			// graph[id] = nil // not need to be nil
		} else {
			graphLines = append(graphLines, line)
		}
	}
	fmt.Println("x", allX, "y", allY)
	fmt.Println("z", binaryToDecimal([]rune(allX))+binaryToDecimal([]rune(allY)))
	// vertices[""] = &vertex{}

	return &solution{
		// graph:    graph,
		expectedZ:  binaryToDecimal([]rune(allX)) + binaryToDecimal([]rune(allY)),
		vertices:   vertices,
		graphlines: graphLines,
	}
}

func (v *vertex) process(dict map[string]*vertex) {
	var a, b bool
	if vertex, ok := dict[v.a]; ok {
		a = vertex.res
	}
	if vertex, ok := dict[v.b]; ok {
		b = vertex.res
	}
	switch v.op {
	case "AND":
		v.res = a && b
	case "XOR":
		v.res = a != b
	case "OR":
		v.res = a || b
	default:
	}
}

type (
	vertex struct {
		a, b, id, op string
		res          bool
	}
	solution struct {
		graphlines    []string
		expectedZ     int
		graph         map[string][]string
		inDegree      map[string]int
		ans           int
		vertices      map[string]*vertex
		sorted, queue []string
	}
)

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
