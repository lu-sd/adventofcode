package main

import (
	"adventofcode/utils"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

var ptoRune = map[utils.Pt]rune{
	{C: 0, R: 1}:  'v',
	{C: -1, R: 0}: '<',
	{C: 0, R: -1}: '^',
	{C: 1, R: 0}:  '>',
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

func (s *solution) find_paths(p utils.Pt, destRune rune, grid utils.Grid[rune]) {
	if !grid.IsInside(p) || s.seen[p] || grid.Get(p) == '#' {
		return
	}
	if grid.Get(p) == destRune {
		if s.min == 0 || len(s.path) <= s.min {
			s.min = len(s.path)
			s.pathsByLen[s.min] = append(s.pathsByLen[s.min], string(s.path))
		}
		return
	}

	s.seen[p] = true
	for _, dir := range utils.Dir4 {
		nextP := p.PMove(dir)
		s.path = append(s.path, ptoRune[dir])
		s.find_paths(nextP, destRune, grid)
		s.path = s.path[:len(s.path)-1]
	}
	s.seen[p] = false
}

func (s *solution) buildMemo(line string, grid utils.Grid[rune], rToP map[rune]utils.Pt) {
	for _, start := range line {
		for _, end := range line {
			pairs := pair{start, end}
			if start == end {
				s.memo[pairs] = []string{""}
			}
			// reinitial the status
			s.min, s.pathsByLen = 0, map[int][]string{}
			s.find_paths(rToP[start], end, grid)
			s.memo[pairs] = s.pathsByLen[s.min]
		}
	}
}

func (s *solution) aggred(str string) []string {
	str = "A" + str
	res, path := []string{}, []string{}
	var dfs func(int)
	dfs = func(level int) {
		if level == len(str)-1 {
			res = append(res, strings.Join(path, "A")+"A")
			return
		}
		start, end := rune(str[level]), rune(str[level+1])
		for _, pattern := range s.memo[pair{start, end}] {
			path = append(path, pattern)
			dfs(level + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return res
}

func (s *solution) dfsmin(str string, blink, target int) int {
	if blink == target {
		return len(str)
	}
	id := step{str, blink}
	if v, ok := s.blinkMemo[id]; ok {
		return v
	}
	str = "A" + str // for range not for build path
	// expand
	ans := 0
	for i := range len(str) - 1 {
		start, end := rune(str[i]), rune(str[i+1])

		res := 0
		for _, path := range s.memo[pair{start, end}] {
			temp := s.dfsmin(path+"A", blink+1, target)
			if res == 0 || temp < res {
				res = temp
			}
		}

		ans += res
	}
	s.blinkMemo[id] = ans
	return ans
}

func (s *solution) run2() {
	// build grid and r to pt map
	s.numG.Array, s.runeToPtNum = s.buildGrid(4, 3, "789456123#0A")
	s.dirG.Array, s.runeToPtDir = s.buildGrid(2, 3, "#^A<v>")
	// for _, row := range s.numG.Array {
	// 	fmt.Println(string(row))
	// }

	// build memo
	s.buildMemo("7894561230A", s.numG, s.runeToPtNum)
	s.buildMemo("^A<v>", s.dirG, s.runeToPtDir)

	// for each line build paths
	// for _, r1Line := range s.input {
	// 	min := 0
	// 	r1 := s.aggred(r1Line)
	//
	// 	for _, line := range r1 {
	// 		// fmt.Println("r2", line)
	// 		r2 := s.aggred(line)
	// 		for _, line := range r2 {
	// 			r3 := s.aggred(line)
	// 			for _, v := range r3 {
	// 				if min == 0 || len(v) < min {
	// 					min = len(v)
	// 				}
	// 			}
	// 		}
	// 	}
	// 	s.ans += min * utils.IntsFromString(r1Line)[0]
	// }
	fmt.Println("original", s.ans)
	res2 := 0
	for _, r1Line := range s.input {
		res := s.dfsmin(r1Line, 0, 26)
		res *= utils.IntsFromString(r1Line)[0]
		res2 += res
	}
	fmt.Println("dfs res", res2)
	// for k, v := range s.memo {
	// 	fmt.Printf("%cto%c \n", k[0], k[1])
	// 	for _, line := range v {
	// 		fmt.Print(string(line), ",")
	// 	}
	// 	fmt.Println("")
	// }
}

func (s *solution) res() int {
	return s.ans
}

type (
	pair [2]rune
	step struct {
		str   string
		blink int
	}
	solution struct {
		input                    []string
		r1, r2, r3               []string
		numG, dirG               utils.Grid[rune]
		seen                     map[utils.Pt]bool
		ans                      int
		min                      int
		pathsByLen               map[int][]string
		path                     []rune
		runeToPtNum, runeToPtDir map[rune]utils.Pt
		memo                     map[pair][]string // start end// shortest paths
		blinkMemo                map[step]int
	}
)

func (s *solution) buildGrid(nrow, ncol int, str string) ([][]rune, map[rune]utils.Pt) {
	grid := make([][]rune, nrow)
	rtoP := map[rune]utils.Pt{}
	count := 0
	for r := range grid {
		grid[r] = make([]rune, ncol)
		for c := range grid[r] {
			R := rune(str[count])
			grid[r][c] = R
			// record run pt map
			rtoP[R] = utils.Pt{C: c, R: r}
			count++
		}
	}
	return grid, rtoP
}

func buildSolution(r io.Reader) *solution {
	lines, err := utils.LinesFromReader(r)
	if err != nil {
		log.Fatalf("could not read input: %v %v", lines, err)
	}

	return &solution{
		input:     lines,
		seen:      map[utils.Pt]bool{},
		memo:      map[pair][]string{},
		blinkMemo: map[step]int{},
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
