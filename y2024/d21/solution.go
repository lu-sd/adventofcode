package main

import (
	"adventofcode/utils"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"time"
)

var ptoRmap = map[utils.Pt]rune{
	{C: 0, R: 1}:  'v',
	{C: -1, R: 0}: '<',
	{C: 0, R: -1}: '^',
	{C: 1, R: 0}:  '>',
}

type pair [2]rune

type solution struct {
	input      []string
	numG, dirG utils.Grid[rune]
	seen       map[utils.Pt]bool
	// nPath        []utils.Pt
	// nRes         [4][][]utils.Pt
	ans int
	// Nshorted     [4]int
	shortest       int
	pathsBylen     map[int][][]rune
	path           []rune
	rpMapN, rpMapD map[rune]utils.Pt
	nPaths         map[pair][][]rune // start end// shortest paths
}

// func (s *solution) cleanNres() {
// 	for i, paths := range s.nRes {
// 		var filtedPath [][]utils.Pt
// 		for _, path := range paths {
// 			if len(path) == s.Nshorted[i] {
// 				filtedPath = append(filtedPath, path)
// 			}
// 		}
// 		s.nRes[i] = filtedPath
// 	}
// }

// func (s *solution) run1() {
// 	for _, line := range s.input {
// 		line = "A" + line
// 		fmt.Println(line)
// 		for i := range line {
// 			if i == len(line)-1 {
// 				break
// 			}
// 			startByte, endByte := line[i], line[i+1]
// 			s.numDfs(s.numMap[rune(startByte)], s.numMap[rune(endByte)], i)
// 		}
// 	}
// 	fmt.Println(s.Nshorted)
// 	s.cleanNres()
// 	var res []string
// 	var buildsfs func(int, string)
// 	buildsfs = func(level int, cur string) {
// 		if level == 4 {
// 			res = append(res, cur)
// 			return
// 		}
// 		for _, pslice := range s.nRes[level] {
// 			rslice := make([]rune, len(pslice))
// 			for i, p := range pslice {
// 				rslice[i] = ptoRmap[p]
// 			}
// 			buildsfs(level+1, cur+string(rslice)+"A")
// 		}
// 	}
// 	buildsfs(0, "")
//
// 	fmt.Println(res)
// }

// func (s *solution) numDfs(p, dest utils.Pt, n int) {
// 	if !s.numG.IsInside(p) || s.seen[p] || s.numG.Get(p) == '#' {
// 		return
// 	}
// 	if p == dest {
// 		if s.Nshorted[n] == 0 || len(s.nPath) <= s.Nshorted[n] {
// 			s.Nshorted[n] = len(s.nPath)
// 			s.nRes[n] = append(s.nRes[n], slices.Clone(s.nPath))
// 		}
// 		return
// 	}
//
// 	s.seen[p] = true
// 	for _, dir := range utils.Dir4 {
// 		nextP := p.PMove(dir)
// 		s.nPath = append(s.nPath, dir)
// 		s.numDfs(nextP, dest, n)
// 		s.nPath = s.nPath[:len(s.nPath)-1]
// 	}
// 	s.seen[p] = false
// }

// update numberPath dict
func (s *solution) all_num_paths(p utils.Pt, destRune rune) {
	if !s.numG.IsInside(p) || s.seen[p] || s.numG.Get(p) == '#' {
		return
	}
	if s.numG.Get(p) == destRune {
		if s.shortest == 0 || len(s.path) <= s.shortest {
			s.shortest = len(s.path)
			curPaths := s.pathsBylen[s.shortest]
			curPaths = append(curPaths, slices.Clone(s.path))
			s.pathsBylen[s.shortest] = curPaths
		}
		return
	}

	s.seen[p] = true
	for _, dir := range utils.Dir4 {
		nextP := p.PMove(dir)
		s.path = append(s.path, ptoRmap[dir])
		s.all_num_paths(nextP, destRune)
		s.path = s.path[:len(s.path)-1]
	}
	s.seen[p] = false
}

func (s *solution) run2() {
	for _, line := range s.input {
		line = "A" + line
		for i := range len(line) - 1 {
			start, end := rune(line[i]), rune(line[i+1])
			// reinitial the status
			s.shortest, s.pathsBylen = 0, map[int][][]rune{}
			s.all_num_paths(s.rpMapN[start], end)
			// fmt.Printf("%c, %c \n", start, end)
			// fmt.Printf("%#v \n", s.shortestPath)
			s.nPaths[pair{start, end}] = s.pathsBylen[s.shortest]
		}
	}
	fmt.Printf("part2 %#v \n", s.nPaths)
	for k, v := range s.nPaths {
		fmt.Printf("%c, %c \n", k[0], k[1])
		for _, line := range v {
			fmt.Println(string(line))
		}
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

	numpab, dirctPad := "789456123#0A", "#^A<v>"
	numGraph := make([][]rune, 4)
	count := 0
	numMap, DMap := map[rune]utils.Pt{}, map[rune]utils.Pt{}
	for r := range numGraph {
		numGraph[r] = make([]rune, 3)
		for c := range numGraph[r] {
			numGraph[r][c] = rune(numpab[count])
			numMap[rune(numpab[count])] = utils.Pt{C: c, R: r}
			count++
		}
	}
	dirGraph := make([][]rune, 2)
	count = 0
	for r := range dirGraph {
		dirGraph[r] = make([]rune, 3)
		for c := range dirGraph[r] {
			dirGraph[r][c] = rune(dirctPad[count])
			DMap[rune(dirctPad[count])] = utils.Pt{C: c, R: r}
			count++
		}
	}
	return &solution{
		input:  lines,
		rpMapN: numMap,
		rpMapD: DMap,
		seen:   map[utils.Pt]bool{},
		nPaths: map[pair][][]rune{},
		numG: utils.Grid[rune]{
			Array: numGraph,
		},
		dirG: utils.Grid[rune]{
			Array: dirGraph,
		},
	}
}

func part1(r io.Reader) int {
	s := buildSolution(r)
	// s.run1()
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
