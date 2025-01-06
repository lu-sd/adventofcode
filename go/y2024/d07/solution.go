package main

import (
	"adventofcode/utils"
	"fmt"
	"io"
	"log"
	"os"
)

type solution struct {
	ans, ans2 int
	lineN     [][]int
}
type opers func(int, int) int

var oper1 = []opers{
	func(a, b int) int {
		return a + b
	},
	func(a, b int) int {
		return a * b
	},
}

var oper2 = append(oper1, func(a, b int) int {
	fact := 1
	for fact <= b {
		fact *= 10
	}
	return a*fact + b
})

func (s *solution) dfs(level, target, curTotal int, input []int, operation []opers) bool {
	if level == len(input) {
		return curTotal == target
	}
	if curTotal > target {
		return false
	}
	for _, oper := range operation {
		if s.dfs(level+1, target, oper(curTotal, input[level]), input, operation) {
			return true
		}
	}
	return false
}

func (s *solution) isvalid(nums []int) bool {
	target, input := nums[0], nums[1:]
	return s.dfs(1, target, input[0], input, oper1)
}

func (s *solution) isvalid2(nums []int) bool {
	target, input := nums[0], nums[1:]
	return s.dfs(1, target, input[0], input, oper2)
}

func (s *solution) run1() {
	for _, nums := range s.lineN {
		if s.isvalid(nums) {
			s.ans += nums[0]
		}
	}
}

func (s *solution) run2() {
	for _, nums := range s.lineN {
		if s.isvalid2(nums) {
			s.ans2 += nums[0]
		}
	}
}

func (s *solution) res() int {
	return s.ans
}

func (s *solution) res2() int {
	return s.ans2
}

func buildSolution(r io.Reader) *solution {
	lines, err := utils.LinesFromReader(r)
	if err != nil {
		log.Fatalf("could not read input: %v %v", lines, err)
	}
	lineN := make([][]int, 0)
	for _, line := range lines {
		nums := utils.IntsFromString(line)
		lineN = append(lineN, nums)
	}
	return &solution{
		ans:   0,
		ans2:  0,
		lineN: lineN,
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
	return s.res2()
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
