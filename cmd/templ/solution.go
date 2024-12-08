package main

import (
	"adventofcode/utils"
	"fmt"
	"io"
	"log"
	"os"
)

type solution struct {
	ans int
}

func (s *solution) run1() {
}

func (s *solution) run2() {
}

func (s *solution) res() int {
	return 0
}

func buildSolution(r io.Reader) solution {
	lines, err := utils.LinesFromReader(r)
	if err != nil {
		log.Fatalf("could not read input: %v %v", lines, err)
	}

	return solution{}
}

func Part1(r io.Reader) int {
	solution := buildSolution(r)
	solution.run1()
	return solution.res()
}

func Part2(r io.Reader) int {
	solution := buildSolution(r)
	solution.run2()
	return solution.res()
}

func main() {
	arg := os.Args[1]
	fmt.Println("Running part", arg)
	switch arg {
	case "1":
		fmt.Println("p1 res 🙆-> ", Part1(os.Stdin))
	case "2":
		fmt.Println("p2 res 🙆-> ", Part2(os.Stdin))
	}
}
