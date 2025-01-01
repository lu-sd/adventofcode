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
	nums      [][]int
}

func (s *solution) isSafe(list []int) bool {
	diff := list[1] - list[0]
	for i := 1; i < len(list); i++ {
		curDif := list[i] - list[i-1]
		if s.invalid(diff, curDif) {
			return false
		}
	}
	return true
}

func (s *solution) invalid(diff, curDif int) bool {
	return diff*curDif <= 0 || utils.Abs(curDif) > 3
}

func (s *solution) run1() {
	for _, numL := range s.nums {
		if s.isSafe(numL) {
			s.ans++
		}
	}
}

func (s *solution) run2() {
	for _, numL := range s.nums {
		if s.isSafe(numL) {
			s.ans2++
		} else {
			for skip := 0; skip < len(numL); skip++ {
				newNumL := []int{}
				for i := 0; i < len(numL); i++ {
					if i == skip {
						continue
					}
					newNumL = append(newNumL, numL[i])
				}
				// newNumL := append(numL[:skip], numL[skip+1:]...) wrong
				if s.isSafe(newNumL) {
					s.ans2++
					break
				}
			}
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
	nums := [][]int{}
	for _, line := range lines {
		numList := utils.IntsFromString(line)
		nums = append(nums, numList)
	}
	return &solution{
		ans:  0,
		ans2: 0,
		nums: nums,
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
	return s.ans2
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
