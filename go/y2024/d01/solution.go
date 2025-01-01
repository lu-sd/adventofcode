package main

import (
	"adventofcode/utils"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

type solution struct {
	ans  int
	arr1 []int
	arr2 []int
	fre  map[int]int
}

func (s *solution) run1() {
	sort.Ints(s.arr1)
	sort.Ints(s.arr2)
	for i := 0; i < len(s.arr1); i++ {
		s.ans += utils.Abs(s.arr1[i] - s.arr2[i])
	}
}

func (s *solution) run2() int {
	res := 0
	for _, num := range s.arr1 {
		res += num * s.fre[num]
	}
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
	arr1 := []int{}
	arr2 := []int{}
	for _, line := range lines {
		nums := utils.IntsFromString(line)
		arr1 = append(arr1, nums[0])
		arr2 = append(arr2, nums[1])
	}

	fre := map[int]int{}
	for _, num := range arr2 {
		fre[num]++
	}

	return &solution{
		ans:  0,
		arr1: arr1,
		arr2: arr2,
		fre:  fre,
	}
}

func part1(r io.Reader) int {
	s := buildSolution(r)
	s.run1()
	return s.res()
}

func part2(r io.Reader) int {
	s := buildSolution(r)
	return s.run2()
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
